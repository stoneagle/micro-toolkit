package controllers

import (
	"encoding/json"
	"strconv"
	"toolkit/backend/common"
	"toolkit/backend/models"
	"toolkit/backend/services"

	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
)

type AutoBuild struct {
	Base
}

func NewAutoBuild(engine *xorm.Engine) *AutoBuild {
	autobuild := &AutoBuild{}
	autobuild.Init()
	return autobuild
}

func (c *AutoBuild) Router(router *gin.RouterGroup) {
	autobuilds := router.Group("autobuild")
	autobuilds.GET("", c.List)
	autobuilds.POST("", c.Create)
	autobuilds.DELETE("/:id", c.Delete)
	autobuilds.POST("/:id/cms", c.Cms)
	autobuilds.POST("/:id/mqtt", c.Mqtt)
	autobuilds.POST("/:id/callback", c.Callback)
	autobuilds.POST("/:id/upgrade", c.Upgrade)
	autobuilds.POST("/:id/album", c.Album)
}

func (c *AutoBuild) List(ctx *gin.Context) {
	engine := common.GetEngine(c.Config.Storybox.Toolkit.Database.Name)
	svc := services.NewAutoBuild(engine)
	autobuilds, err := svc.GetList()
	if err != nil {
		c.ErrorBusiness(ctx, common.ErrorMysql, "autobuild get error", err)
		return
	}
	c.Success(ctx, autobuilds)
}

func (c *AutoBuild) Create(ctx *gin.Context) {
	var autobuild models.AutoBuild
	if err := ctx.ShouldBindJSON(&autobuild); err != nil {
		c.ErrorBusiness(ctx, common.ErrorParams, "params error: ", err)
		return
	}

	engine := common.GetEngine(c.Config.Storybox.Toolkit.Database.Name)
	_, err := engine.Insert(&autobuild)
	if err != nil {
		c.ErrorBusiness(ctx, common.ErrorMysql, "autobuild insert error", err)
		return
	}

	c.Success(ctx, autobuild)
}

func (c *AutoBuild) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.ErrorBusiness(ctx, common.ErrorParams, "id params error", err)
		return
	}

	engine := common.GetEngine(c.Config.Storybox.Toolkit.Database.Name)
	var autobuild models.AutoBuild
	_, err = engine.Id(id).Get(&autobuild)
	if err != nil {
		c.ErrorBusiness(ctx, common.ErrorMysql, "autobuild get error", err)
		return
	}

	_, err = engine.Id(id).Delete(&autobuild)
	if err != nil {
		c.ErrorBusiness(ctx, common.ErrorMysql, "autobuild delete error", err)
		return
	}

	c.Success(ctx, id)
}

func (c *AutoBuild) Cms(ctx *gin.Context) {
}

func (c *AutoBuild) Mqtt(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.ErrorBusiness(ctx, common.ErrorParams, "id params error", err)
		return
	}

	engineTK := common.GetEngine(c.Config.Storybox.Toolkit.Database.Name)
	engineMqtt := common.GetEngine(c.Config.Storybox.Mqtt.Database.Name)

	svc := services.NewPushChannel(engineTK, engineMqtt)
	err = svc.Add(id, c.Config.Storybox.Mqtt.Params)
	if err != nil {
		c.ErrorBusiness(ctx, common.ErrorMysql, "mqtt配置失败", err)
		return
	}
	c.Success(ctx, id)
}

func (c *AutoBuild) Callback(ctx *gin.Context) {
	var autobuild models.AutoBuild
	if err := ctx.ShouldBindJSON(&autobuild); err != nil {
		c.ErrorBusiness(ctx, common.ErrorParams, "params error: ", err)
		return
	}

	if autobuild.Id == 0 || autobuild.Callback == "" {
		c.ErrorBusiness(ctx, common.ErrorParams, "id or callback params not exist", nil)
		return
	}

	engineTK := common.GetEngine(c.Config.Storybox.Toolkit.Database.Name)
	engineCB := common.GetEngine(c.Config.Storybox.Callback.Database.Name)

	templateConfig := c.Config.Storybox.Callback.Config
	var templateMap map[string][]models.CallbackTemplate
	err := json.Unmarshal([]byte(templateConfig), &templateMap)
	if err != nil {
		c.ErrorBusiness(ctx, common.ErrorParams, "callback config can not exec json unmarshal", err)
		return
	}

	templateSlice, ok := templateMap[autobuild.Callback]
	if !ok {
		c.ErrorBusiness(ctx, common.ErrorParams, "type relate config not exist", nil)
		return
	}

	svc := services.NewCallbackConfig(engineTK, engineCB)
	err = svc.Add(int(autobuild.Id), templateSlice, autobuild.Callback)
	if err != nil {
		c.ErrorBusiness(ctx, common.ErrorMysql, "callback配置失败", err)
		return
	}

	c.Success(ctx, autobuild.Id)
}

func (c *AutoBuild) Upgrade(ctx *gin.Context) {
	var autobuild models.AutoBuild
	if err := ctx.ShouldBindJSON(&autobuild); err != nil {
		c.ErrorBusiness(ctx, common.ErrorParams, "params error: ", err)
		return
	}

	if autobuild.Id == 0 || autobuild.UpgradeName == "" || autobuild.UpgradeVcode == 0 {
		c.ErrorBusiness(ctx, common.ErrorParams, "id, name or vcode params can not be empty", nil)
		return
	}

	engineTK := common.GetEngine(c.Config.Storybox.Toolkit.Database.Name)
	engineUP := common.GetEngine(c.Config.Storybox.Upgrade.Database.Name)
	svc := services.NewUpUpdate(engineTK, engineUP)
	err := svc.Add(int(autobuild.Id), int(autobuild.UpgradeVcode), autobuild.UpgradeName, autobuild.UpgradeVname)
	if err != nil {
		c.ErrorBusiness(ctx, common.ErrorMysql, "upgrade配置失败", err)
		return
	}
	c.Success(ctx, int(autobuild.Id))
}

func (c *AutoBuild) Album(ctx *gin.Context) {
	var autobuild models.AutoBuild
	if err := ctx.ShouldBindJSON(&autobuild); err != nil {
		c.ErrorBusiness(ctx, common.ErrorParams, "params error: ", err)
		return
	}
	if autobuild.Id == 0 || autobuild.AlbumList == "" {
		c.ErrorBusiness(ctx, common.ErrorParams, "id and albumList params can not be empty", nil)
		return
	}
	engineTK := common.GetEngine(c.Config.Storybox.Toolkit.Database.Name)
	engineAL := common.GetEngine(c.Config.Storybox.Album.Database.Name)
	svc := services.NewCmsPresetAlbums(engineTK, engineAL)
	err := svc.Add(int(autobuild.Id), autobuild.AlbumList)
	if err != nil {
		c.ErrorBusiness(ctx, common.ErrorMysql, "albumlist配置失败", err)
		return
	}
	c.Success(ctx, int(autobuild.Id))
}
