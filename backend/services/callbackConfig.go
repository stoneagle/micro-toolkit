package services

import (
	"time"
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

func (u *CallbackConfig) List(appId string) (callbacks []models.CallbackConfig, err error) {
	callbacks = make([]models.CallbackConfig, 0)
	err = u.engineCB.Where("app_id = ?", appId).Find(&callbacks)
	return
}

func (s *CallbackConfig) Delete(appId string) (err error) {
	session := s.engineCB.NewSession()
	defer session.Close()
	lists, err := s.List(appId)
	if err != nil {
		session.Rollback()
		return
	}
	for _, one := range lists {
		_, err = session.Id(one.General.Id).Delete(&one)
		if err != nil {
			session.Rollback()
			return
		}
	}
	err = session.Commit()
	return
}

func (u *CallbackConfig) Add(autobuildId int, templateSlice []models.CallbackTemplate, ctype string) (err error) {
	sessionTK := u.engineTK.NewSession()
	sessionCB := u.engineCB.NewSession()
	defer sessionTK.Close()
	defer sessionCB.Close()
	err = sessionTK.Begin()
	if err != nil {
		return
	}
	err = sessionCB.Begin()
	if err != nil {
		return
	}

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
			State:         1,
			CallbackState: template.CallbackState,
			CallbackUrl:   template.Url,
			Action:        template.Action,
		}
		callback.General.UpdatedAt = int(time.Now().Unix())
		callback.General.CreatedAt = int(time.Now().Unix())
		_, err = sessionCB.Insert(&callback)
		if err != nil {
			sessionTK.Rollback()
			sessionCB.Rollback()
			return err
		}
	}

	updateAutobuild := models.AutoBuild{
		Callback: ctype,
	}
	_, err = sessionTK.Where("id = ?", autobuild.Id).Update(&updateAutobuild)
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
