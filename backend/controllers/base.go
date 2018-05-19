package controllers

import (
	"fmt"
	"net/http"
	"toolkit/backend/common"

	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
)

type Base struct {
	Config common.Conf
	Engine *xorm.Engine
}

func (b *Base) Init() {
	b.Config = *common.GetConfig()
}

func (b *Base) Redirect(ctx *gin.Context, uri string) {
	ctx.Redirect(http.StatusFound, uri)
}

func (b *Base) Success(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": common.ErrorOk,
		"data": data,
		"desc": "",
	})
}

func (b *Base) ErrorBusiness(ctx *gin.Context, code common.ErrorCode, desc string, err error) {
	fmt.Println("%v\r\n", err.Error())
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": struct{}{},
		"desc": desc,
	})
}

func (b *Base) ErrorServer(ctx *gin.Context, desc string) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": common.ErrorServer,
		"data": struct{}{},
		"desc": desc,
	})
}
