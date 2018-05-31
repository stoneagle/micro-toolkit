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
	AutoBuildSvc *services.AutoBuild
	MqttSvc      *services.PushChannel
	CallbackSvc  *services.CallbackConfig
	UpgradeSvc   *services.UpUpdate
	AlbumSvc     *services.CmsPresetAlbums
}

func NewAutoBuild(engine *xorm.Engine) *AutoBuild {
	autobuild := &AutoBuild{}
	autobuild.Init()
	engineTK := common.GetEngine(autobuild.Config.Storybox.Toolkit.Database.Name)
	engineMqtt := common.GetEngine(autobuild.Config.Storybox.Mqtt.Database.Name)
	engineCB := common.GetEngine(autobuild.Config.Storybox.Callback.Database.Name)
	engineUP := common.GetEngine(autobuild.Config.Storybox.Upgrade.Database.Name)
	engineAL := common.GetEngine(autobuild.Config.Storybox.Album.Database.Name)

	autobuild.AutoBuildSvc = services.NewAutoBuild(engineTK)
	autobuild.MqttSvc = services.NewPushChannel(engineTK, engineMqtt)
	autobuild.CallbackSvc = services.NewCallbackConfig(engineTK, engineCB)
	autobuild.UpgradeSvc = services.NewUpUpdate(engineTK, engineUP)
	autobuild.AlbumSvc = services.NewCmsPresetAlbums(engineTK, engineAL)
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
	autobuilds.GET("/:id/mqtt", InitAutobuild(), c.MqttList)
	autobuilds.GET("/:id/callback", InitAutobuild(), c.CallbackList)
	autobuilds.GET("/:id/upgrade", InitAutobuild(), c.UpgradeList)
	autobuilds.GET("/:id/album", InitAutobuild(), c.AlbumList)
}

func InitAutobuild() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idStr := ctx.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			common.ResponseErrorBusiness(ctx, common.ErrorParams, "id params error", err)
		}

		engineTK := common.GetEngine(common.GetConfig().Storybox.Toolkit.Database.Name)
		autoBuildSvc := services.NewAutoBuild(engineTK)
		autobuild, err := autoBuildSvc.GetOne(id)
		if err != nil {
			common.ResponseErrorBusiness(ctx, common.ErrorMysql, "get autobuild error", err)
		}

		ctx.Set("autobuild", autobuild)
		return
	}
}

func (c *AutoBuild) List(ctx *gin.Context) {
	autobuilds, err := c.AutoBuildSvc.GetList()
	if err != nil {
		common.ResponseErrorBusiness(ctx, common.ErrorMysql, "autobuild get error", err)
		return
	}
	common.ResponseSuccess(ctx, autobuilds)
}

func (c *AutoBuild) Create(ctx *gin.Context) {
	var autobuild models.AutoBuild
	if err := ctx.ShouldBindJSON(&autobuild); err != nil {
		common.ResponseErrorBusiness(ctx, common.ErrorParams, "params error: ", err)
		return
	}

	err := c.AutoBuildSvc.Add(&autobuild)
	if err != nil {
		common.ResponseErrorBusiness(ctx, common.ErrorMysql, "autobuild insert error", err)
		return
	}

	common.ResponseSuccess(ctx, autobuild)
}

func (c *AutoBuild) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		common.ResponseErrorBusiness(ctx, common.ErrorParams, "id params error", err)
		return
	}

	err = c.AutoBuildSvc.Delete(id)
	if err != nil {
		common.ResponseErrorBusiness(ctx, common.ErrorMysql, "autobuild delete error", err)
		return
	}

	common.ResponseSuccess(ctx, id)
}

func (c *AutoBuild) Cms(ctx *gin.Context) {
}

func (c *AutoBuild) Mqtt(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		common.ResponseErrorBusiness(ctx, common.ErrorParams, "id params error", err)
		return
	}

	err = c.MqttSvc.Add(id, c.Config.Storybox.Mqtt.Params)
	if err != nil {
		common.ResponseErrorBusiness(ctx, common.ErrorMysql, "mqtt配置失败", err)
		return
	}
	common.ResponseSuccess(ctx, id)
}

func (c *AutoBuild) Callback(ctx *gin.Context) {
	var autobuild models.AutoBuild
	if err := ctx.ShouldBindJSON(&autobuild); err != nil {
		common.ResponseErrorBusiness(ctx, common.ErrorParams, "params error: ", err)
		return
	}

	if autobuild.Id == 0 || autobuild.Callback == "" {
		common.ResponseErrorBusiness(ctx, common.ErrorParams, "id or callback params not exist", nil)
		return
	}

	templateConfig := c.Config.Storybox.Callback.Config
	var templateMap map[string][]models.CallbackTemplate
	err := json.Unmarshal([]byte(templateConfig), &templateMap)
	if err != nil {
		common.ResponseErrorBusiness(ctx, common.ErrorParams, "callback config can not exec json unmarshal", err)
		return
	}

	templateSlice, ok := templateMap[autobuild.Callback]
	if !ok {
		common.ResponseErrorBusiness(ctx, common.ErrorParams, "type relate config not exist", nil)
		return
	}

	err = c.CallbackSvc.Add(autobuild.Id, templateSlice, autobuild.Callback)
	if err != nil {
		common.ResponseErrorBusiness(ctx, common.ErrorMysql, "callback配置失败", err)
		return
	}

	common.ResponseSuccess(ctx, autobuild.Id)
}

func (c *AutoBuild) Upgrade(ctx *gin.Context) {
	var autobuild models.AutoBuild
	if err := ctx.ShouldBindJSON(&autobuild); err != nil {
		common.ResponseErrorBusiness(ctx, common.ErrorParams, "params error: ", err)
		return
	}

	if autobuild.Id == 0 || autobuild.UpgradeName == "" || autobuild.UpgradeVcode == 0 {
		common.ResponseErrorBusiness(ctx, common.ErrorParams, "id, name or vcode params can not be empty", nil)
		return
	}

	err := c.UpgradeSvc.Add(autobuild.Id, autobuild.UpgradeVcode, autobuild.UpgradeName, autobuild.UpgradeVname)
	if err != nil {
		common.ResponseErrorBusiness(ctx, common.ErrorMysql, "upgrade配置失败", err)
		return
	}
	common.ResponseSuccess(ctx, autobuild.Id)
}

func (c *AutoBuild) Album(ctx *gin.Context) {
	var autobuild models.AutoBuild
	if err := ctx.ShouldBindJSON(&autobuild); err != nil {
		common.ResponseErrorBusiness(ctx, common.ErrorParams, "params error: ", err)
		return
	}

	if autobuild.Id == 0 || autobuild.AlbumList == "" {
		common.ResponseErrorBusiness(ctx, common.ErrorParams, "id and albumList params can not be empty", nil)
		return
	}

	err := c.AlbumSvc.Add(autobuild.Id, autobuild.AlbumList)
	if err != nil {
		common.ResponseErrorBusiness(ctx, common.ErrorMysql, "albumlist配置失败", err)
		return
	}
	common.ResponseSuccess(ctx, autobuild.Id)
}

func (c *AutoBuild) MqttList(ctx *gin.Context) {
	autobuild := ctx.MustGet("autobuild").(models.AutoBuild)

	mqtts, err := c.MqttSvc.List(autobuild.AppId)
	if err != nil {
		common.ResponseErrorBusiness(ctx, common.ErrorMysql, "get pushChannels error", err)
		return
	}

	common.ResponseSuccess(ctx, mqtts)
}

func (c *AutoBuild) CallbackList(ctx *gin.Context) {
	autobuild := ctx.MustGet("autobuild").(models.AutoBuild)

	callbacks, err := c.CallbackSvc.List(autobuild.AppId)
	if err != nil {
		common.ResponseErrorBusiness(ctx, common.ErrorMysql, "get callbacks error", err)
		return
	}
	common.ResponseSuccess(ctx, callbacks)
}

func (c *AutoBuild) UpgradeList(ctx *gin.Context) {
	autobuild := ctx.MustGet("autobuild").(models.AutoBuild)

	upgrades, err := c.UpgradeSvc.List(autobuild.UpgradeName, autobuild.UpgradeVcode)
	if err != nil {
		common.ResponseErrorBusiness(ctx, common.ErrorMysql, "get upgrades error", err)
		return
	}
	common.ResponseSuccess(ctx, upgrades)
}

func (c *AutoBuild) AlbumList(ctx *gin.Context) {
	autobuild := ctx.MustGet("autobuild").(models.AutoBuild)

	albums, err := c.AlbumSvc.List(autobuild.AppId)
	if err != nil {
		common.ResponseErrorBusiness(ctx, common.ErrorMysql, "get album error", err)
		return
	}
	common.ResponseSuccess(ctx, albums)
}
