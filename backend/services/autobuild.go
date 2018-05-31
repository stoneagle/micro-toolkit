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

func (m *AutoBuild) UpdateByMap(id int, autobuild map[string]interface{}) (err error) {
	_, err = m.engine.Table(new(models.AutoBuild)).Id(id).Update(&autobuild)
	return
}

func (m *AutoBuild) GetOne(id int) (autobuild models.AutoBuild, err error) {
	autobuild = models.AutoBuild{}
	_, err = m.engine.Where("id = ?", id).Get(&autobuild)
	return
}

func (m *AutoBuild) GetList() (autobuilds []models.AutoBuild, err error) {
	autobuilds = make([]models.AutoBuild, 0)
	err = m.engine.Find(&autobuilds)
	return
}

func (m *AutoBuild) Add(autobuild *models.AutoBuild) (err error) {
	_, err = m.engine.Insert(autobuild)
	return
}

func (m *AutoBuild) Delete(id int) (err error) {
	var autobuild models.AutoBuild
	_, err = m.engine.Id(id).Get(&autobuild)
	if err == nil {
		_, err = m.engine.Id(id).Delete(&autobuild)
	}
	return
}
