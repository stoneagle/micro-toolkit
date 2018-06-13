package controllers

import (
	"strconv"
	"toolkit/backend/common"
	"toolkit/backend/models"
	"toolkit/backend/services"

	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
)

type Scan struct {
	Base
	ScanSvc *services.DeviceScan
}

func NewScan(engine *xorm.Engine) *Scan {
	scan := &Scan{}
	scan.Init()
	engine = common.GetEngine(scan.Config.Storybox.Rdevice.Database.Name)

	scan.ScanSvc = services.NewDeviceScan(engine)
	return scan
}

func (c *Scan) Router(router *gin.RouterGroup) {
	scans := router.Group("scan")
	scans.GET("/:appId", c.ListByApp)
	scans.POST("", c.Create)
	scans.PUT("", c.Update)
	scans.DELETE("/:id", InitScan(c.ScanSvc), c.Delete)
}

func InitScan(svc *services.DeviceScan) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idStr := ctx.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			common.ResponseErrorBusiness(ctx, common.ErrorParams, "id params error", err)
		}

		scan, err := svc.One(id)
		if err != nil {
			common.ResponseErrorBusiness(ctx, common.ErrorMysql, "get scan error", err)
		}

		ctx.Set("scan", scan)
		return
	}
}

func (c *Scan) ListByApp(ctx *gin.Context) {
	appId := ctx.Param("appId")
	if appId == "" {
		common.ResponseErrorBusiness(ctx, common.ErrorParams, "appId params not exist", nil)
		return
	}
	scans, err := c.ScanSvc.List(appId)
	if err != nil {
		common.ResponseErrorBusiness(ctx, common.ErrorMysql, "scan get error", err)
		return
	}
	common.ResponseSuccess(ctx, scans)
}

func (c *Scan) Create(ctx *gin.Context) {
	var scan models.DeviceScan
	if err := ctx.ShouldBindJSON(&scan); err != nil {
		common.ResponseErrorBusiness(ctx, common.ErrorParams, "params error: ", err)
		return
	}

	err := c.ScanSvc.Add(&scan)
	if err != nil {
		common.ResponseErrorBusiness(ctx, common.ErrorMysql, "scan insert error", err)
		return
	}

	common.ResponseSuccess(ctx, scan)
}

func (c *Scan) Update(ctx *gin.Context) {
	var scan models.DeviceScan
	if err := ctx.ShouldBindJSON(&scan); err != nil {
		common.ResponseErrorBusiness(ctx, common.ErrorParams, "params error: ", err)
		return
	}

	err := c.ScanSvc.Update(scan.Id, &scan)
	if err != nil {
		common.ResponseErrorBusiness(ctx, common.ErrorMysql, "scan update error", err)
		return
	}

	common.ResponseSuccess(ctx, scan)
}

func (c *Scan) Delete(ctx *gin.Context) {
	scan := ctx.MustGet("scan").(models.DeviceScan)
	err := c.ScanSvc.Delete(scan.Id)
	if err != nil {
		common.ResponseErrorBusiness(ctx, common.ErrorMysql, "scan delete error", err)
		return
	}

	common.ResponseSuccess(ctx, scan.Id)
}
