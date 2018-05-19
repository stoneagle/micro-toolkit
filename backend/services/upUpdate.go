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
	sessionUP := u.engineUP.NewSession()
	defer sessionUP.Close()

	update := models.UpUpdate{
		Name:  name,
		Vcode: uint(vcode),
		Vname: vname,
	}
	_, err = sessionUP.Insert(&update)
	if err != nil {
		sessionUP.Rollback()
		return err
	}
	err = sessionUP.Commit()
	if err != nil {
		sessionUP.Rollback()
		return err
	}
	return nil
}
