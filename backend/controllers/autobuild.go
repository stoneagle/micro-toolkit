package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
)

type AutoBuild struct {
	Base
}

func NewAutoBuild(engine *xorm.Engine) *AutoBuild {
	autobuild := &AutoBuild{}
	autobuild.Init()
	autobuild.Engine = engine
	return autobuild
}

func (c *AutoBuild) Router(router *gin.RouterGroup) {
	autobuilds := router.Group("autobuild")
	autobuilds.GET("", c.List)
	autobuilds.POST("", c.Create)
	autobuilds.PUT("", c.Update)
	autobuilds.POST("/cms", c.Cms)
	autobuilds.POST("/mqtt", c.Mqtt)
	autobuilds.POST("/callback", c.Callback)
	autobuilds.POST("/upgrade", c.Upgrade)
	autobuilds.POST("/album", c.Album)
}

func (c *AutoBuild) List(ctx *gin.Context) {
}

func (c *AutoBuild) Create(ctx *gin.Context) {
}

func (c *AutoBuild) Update(ctx *gin.Context) {
}

func (c *AutoBuild) Cms(ctx *gin.Context) {
}

func (c *AutoBuild) Mqtt(ctx *gin.Context) {
}

func (c *AutoBuild) Callback(ctx *gin.Context) {
}

func (c *AutoBuild) Upgrade(ctx *gin.Context) {
}

func (c *AutoBuild) Album(ctx *gin.Context) {
}
