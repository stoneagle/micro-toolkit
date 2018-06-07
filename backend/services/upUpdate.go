package services

import (
	"time"
	"toolkit/backend/models"

	"github.com/go-xorm/xorm"
)

type UpUpdate struct {
	engineTK *xorm.Engine
	engineUP *xorm.Engine
}

func NewUpUpdate(tk, up *xorm.Engine) *UpUpdate {
	return &UpUpdate{
		engineTK: tk,
		engineUP: up,
	}
}

func (u *UpUpdate) List(name string, vcode int) (updates []models.UpUpdate, err error) {
	updates = make([]models.UpUpdate, 0)
	err = u.engineUP.Unscoped().Where("name = ?", name).And("vcode = ?", vcode).Find(&updates)
	return
}

func (s *UpUpdate) Delete(name string, vcode int) (err error) {
	session := s.engineUP.NewSession()
	defer session.Close()
	lists, err := s.List(name, vcode)
	if err != nil {
		session.Rollback()
		return
	}
	for _, one := range lists {
		_, err = session.Unscoped().Id(one.General.Id).Delete(&one)
		if err != nil {
			session.Rollback()
			return
		}
	}
	err = session.Commit()
	return
}

func (u *UpUpdate) Add(autobuildId, vcode int, name, vname string) (err error) {
	sessionTK := u.engineTK.NewSession()
	defer sessionTK.Close()
	sessionUP := u.engineUP.NewSession()
	defer sessionUP.Close()

	autobuild := models.AutoBuild{}
	_, err = sessionTK.Where("id = ?", autobuildId).Get(&autobuild)
	if err != nil {
		sessionTK.Rollback()
		sessionUP.Rollback()
		return err
	}

	update := models.UpUpdate{
		Name:  name,
		Vcode: vcode,
		Vname: vname,
	}
	update.General.UpdatedAt = int(time.Now().Unix())
	update.General.CreatedAt = int(time.Now().Unix())
	_, err = sessionUP.Insert(&update)
	if err != nil {
		sessionTK.Rollback()
		sessionUP.Rollback()
		return err
	}

	updateAutobuild := models.AutoBuild{
		UpgradeName:  name,
		UpgradeVcode: vcode,
		UpgradeVname: vname,
	}
	_, err = sessionTK.Where("id = ?", autobuild.Id).Update(&updateAutobuild)
	if err != nil {
		sessionTK.Rollback()
		sessionUP.Rollback()
		return err
	}

	err = sessionTK.Commit()
	if err != nil {
		sessionTK.Rollback()
		sessionUP.Rollback()
		return err
	}
	err = sessionUP.Commit()
	if err != nil {
		sessionTK.Rollback()
		sessionUP.Rollback()
		return err
	}
	return nil
}
