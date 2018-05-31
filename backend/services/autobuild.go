package services

import (
	"toolkit/backend/models"

	"github.com/go-xorm/xorm"
)

type AutoBuild struct {
	engine *xorm.Engine
}

func NewAutoBuild(engine *xorm.Engine) *AutoBuild {
	return &AutoBuild{
		engine: engine,
	}
}

func (m *AutoBuild) GetList() (autobuilds []models.AutoBuild, err error) {
	autobuilds = make([]models.AutoBuild, 0)
	err = m.engine.Find(&autobuilds)
	return
}
