package controllers

import (
	"toolkit/backend/common"
)

type Base struct {
	Config common.Conf
}

func (b *Base) Init() {
	b.Config = *common.GetConfig()
}
