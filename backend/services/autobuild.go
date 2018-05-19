package services

import "github.com/go-xorm/xorm"

type AutoBuild struct {
	engine *xorm.Engine
}

func NewAutoBuild(engine *xorm.Engine) *AutoBuild {
	return &AutoBuild{
		engine: engine,
	}
}
