package services

import (
	"errors"
	"toolkit/backend/models"

	"github.com/go-xorm/xorm"
)

type PushChannel struct {
	engineTK   *xorm.Engine
	engineMqtt *xorm.Engine
}

func NewPushChannel(tk, mqtt *xorm.Engine) *PushChannel {
	return &PushChannel{
		engineTK:   tk,
		engineMqtt: mqtt,
	}
}

func (p *PushChannel) Add(autobuildId int, paramsJson string) (err error) {
	sessionTK := p.engineTK.NewSession()
	defer sessionTK.Close()
	sessionMqtt := p.engineMqtt.NewSession()
	defer sessionMqtt.Close()

	autobuild := models.AutoBuild{}
	_, err = sessionTK.Where("id = ?", autobuildId).Get(&autobuild)
	if err != nil {
		sessionTK.Rollback()
		sessionMqtt.Rollback()
		return err
	}

	// 如果mqtt状态为true，则检查是否已插入
	if autobuild.Mqtt != 0 {
		tmpMqtt := models.PushChannel{}
		_, err = sessionMqtt.Where("channel_type = ?", autobuild.AppId).And("service = ?", "mqtt").And("production = ?", "storybox").Get(&tmpMqtt)
		if err != nil {
			sessionTK.Rollback()
			sessionMqtt.Rollback()
			return err
		}
		if tmpMqtt.Id > 0 {
			return errors.New("mqtt already set")
		}
	}

	mqtt := models.PushChannel{
		ChannelType: autobuild.AppId,
		Production:  "storybox",
		Service:     "mqtt",
		Params:      paramsJson,
	}
	_, err = sessionMqtt.Insert(&mqtt)
	if err != nil {
		sessionTK.Rollback()
		sessionMqtt.Rollback()
		return err
	}

	autobuild.Mqtt = 1
	_, err = sessionTK.Update(&autobuild)
	if err != nil {
		sessionTK.Rollback()
		sessionMqtt.Rollback()
		return err
	}

	err = sessionTK.Commit()
	if err != nil {
		sessionTK.Rollback()
		sessionMqtt.Rollback()
		return err
	}
	err = sessionMqtt.Commit()
	if err != nil {
		sessionTK.Rollback()
		sessionMqtt.Rollback()
		return err
	}
	return nil
}
