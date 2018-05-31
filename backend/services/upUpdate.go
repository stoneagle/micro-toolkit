package services

import (
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
