package services

import (
	"toolkit/backend/models"

	"github.com/go-xorm/xorm"
)

type CallbackConfig struct {
	engineTK *xorm.Engine
	engineCB *xorm.Engine
}

func NewCallbackConfig(tk, cb *xorm.Engine) *CallbackConfig {
	return &CallbackConfig{
		engineTK: tk,
		engineCB: cb,
	}
}

func (u *CallbackConfig) Add(autobuildId int, templateSlice []models.CallbackTemplate, ctype string) (err error) {
	sessionTK := u.engineTK.NewSession()
	defer sessionTK.Close()
	sessionCB := u.engineCB.NewSession()
	defer sessionCB.Close()

	autobuild := models.AutoBuild{}
	_, err = sessionTK.Where("id = ?", autobuildId).Get(&autobuild)
	if err != nil {
		sessionTK.Rollback()
		sessionCB.Rollback()
		return err
	}

	for _, template := range templateSlice {
		callback := models.CallbackConfig{
			AppId:         autobuild.AppId,
			CallbackState: template.State,
			CallbackUrl:   template.Url,
			Action:        template.Action,
		}
		_, err = sessionCB.Insert(&callback)
		if err != nil {
			sessionTK.Rollback()
			sessionCB.Rollback()
			return err
		}
	}

	autobuild.Callback = ctype
	_, err = sessionTK.Update(&autobuild)
	if err != nil {
		sessionTK.Rollback()
		sessionCB.Rollback()
		return err
	}

	err = sessionTK.Commit()
	if err != nil {
		sessionTK.Rollback()
		sessionCB.Rollback()
		return err
	}
	err = sessionCB.Commit()
	if err != nil {
		sessionTK.Rollback()
		sessionCB.Rollback()
		return err
	}
	return nil
}