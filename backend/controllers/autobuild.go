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
	autobuilds.POST("/:id/cms", c.Cms)
	autobuilds.POST("/:id/mqtt", c.Mqtt)
	autobuilds.POST("/:id/callback", c.Callback)
	autobuilds.POST("/:id/upgrade", c.Upgrade)
	autobuilds.POST("/:id/album", c.Album)
}

func (c *AutoBuild) List(ctx *gin.Context) {
}

func (c *AutoBuild) Create(ctx *gin.Context) {
	appId := ctx.PostForm("appId")
	if appId == "" {
		c.ErrorBusiness(ctx, common.ErrorParams, "appId params not exist", nil)
		return
	}

	engine := common.GetEngine(c.Config.Storybox.Toolkit.Database.Name)
	autobuild := models.AutoBuild{
		AppId: appId,
	}

	_, err := engine.Insert(&autobuild)
	if err != nil {
		c.ErrorBusiness(ctx, common.ErrorMysql, "autobuild insert error", err)
		return
	}

	c.Success(ctx, autobuild.Id)
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
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.ErrorBusiness(ctx, common.ErrorParams, "id params error", err)
		return
	}
	engineTK := common.GetEngine(c.Config.Storybox.Toolkit.Database.Name)
	engineCB := common.GetEngine(c.Config.Storybox.Callback.Database.Name)

	ctype := ctx.PostForm("ctype")
	if ctype == "" {
		c.ErrorBusiness(ctx, common.ErrorParams, "ctype can not be empty", err)
		return
	}

	templateConfig := c.Config.Storybox.Callback.Config
	var templateMap map[string][]models.CallbackTemplate
	err = json.Unmarshal([]byte(templateConfig), &templateMap)
	if err != nil {
		c.ErrorBusiness(ctx, common.ErrorParams, "callback config can not exec json unmarshal", err)
		return
	}

	templateSlice, ok := templateMap[ctype]
	if !ok {
		c.ErrorBusiness(ctx, common.ErrorParams, "type relate config not exist", nil)
		return
	}

	svc := services.NewCallbackConfig(engineTK, engineCB)
	err = svc.Add(id, templateSlice, ctype)
	if err != nil {
		c.ErrorBusiness(ctx, common.ErrorMysql, "callback配置失败", err)
		return
	}

	c.Success(ctx, id)
}

func (c *AutoBuild) Upgrade(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.ErrorBusiness(ctx, common.ErrorParams, "id params error", err)
		return
	}

	name := ctx.PostForm("name")
	vcodeStr := ctx.PostForm("vcode")
	vname := ctx.PostForm("vname")
	if name == "" || vcodeStr == "" {
		c.ErrorBusiness(ctx, common.ErrorParams, "name or vcode params can not be empty", err)
		return
	}
	vcode, err := strconv.Atoi(vcodeStr)
	if err != nil {
		c.ErrorBusiness(ctx, common.ErrorParams, "vcode params error", err)
		return
	}

	engineTK := common.GetEngine(c.Config.Storybox.Toolkit.Database.Name)
	engineUP := common.GetEngine(c.Config.Storybox.Upgrade.Database.Name)
	svc := services.NewUpUpdate(engineTK, engineUP)
	err = svc.Add(id, vcode, name, vname)
	if err != nil {
		c.ErrorBusiness(ctx, common.ErrorMysql, "upgrade配置失败", err)
		return
	}
	c.Success(ctx, id)
}

func (c *AutoBuild) Album(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.ErrorBusiness(ctx, common.ErrorParams, "id params error", err)
		return
	}
	albumList := ctx.PostForm("albumList")
	if albumList == "" {
		c.ErrorBusiness(ctx, common.ErrorParams, "albumList params can not be empty", err)
		return
	}

	engineTK := common.GetEngine(c.Config.Storybox.Toolkit.Database.Name)
	engineAL := common.GetEngine(c.Config.Storybox.Album.Database.Name)
	svc := services.NewCmsPresetAlbums(engineTK, engineAL)
	err = svc.Add(id, albumList)
	if err != nil {
		c.ErrorBusiness(ctx, common.ErrorMysql, "albumlist配置失败", err)
		return
	}
	c.Success(ctx, id)
}
