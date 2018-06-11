package services

import (
	"errors"
	"time"
	"toolkit/backend/models"

	"github.com/go-xorm/xorm"
)

type PushChannel struct {
	engineTK   *xorm.Engine
	engineMqtt *xorm.Engine
	lastfix    string
}

func NewPushChannel(tk, mqtt *xorm.Engine, lastfix string) *PushChannel {
	return &PushChannel{
		engineTK:   tk,
		engineMqtt: mqtt,
		lastfix:    lastfix,
	}
}

func (p *PushChannel) List(channelType string) (pushChannels []models.PushChannel, err error) {
	pushChannels = make([]models.PushChannel, 0)
	err = p.engineMqtt.Where("channel_type = ?", channelType+p.lastfix).And("service = ?", "mqtt").And("production = ?", "storybox").Find(&pushChannels)
	return
}

func (p *PushChannel) Delete(channelType string) (err error) {
	sessionMqtt := p.engineMqtt.NewSession()
	defer sessionMqtt.Close()
	lists, err := p.List(channelType)
	if err != nil {
		sessionMqtt.Rollback()
		return
	}
	for _, one := range lists {
		_, err = sessionMqtt.Id(one.General.Id).Delete(&one)
		if err != nil {
			sessionMqtt.Rollback()
			return
		}
	}
	err = sessionMqtt.Commit()
	return
}

func (p *PushChannel) Add(autobuildId int, paramsJson string) (err error) {
	sessionTK := p.engineTK.NewSession()
	sessionMqtt := p.engineMqtt.NewSession()
	defer sessionTK.Close()
	defer sessionMqtt.Close()
	err = sessionTK.Begin()
	if err != nil {
		return
	}
	err = sessionMqtt.Begin()
	if err != nil {
		return
	}

	autobuild := models.AutoBuild{}
	_, err = sessionTK.Where("id = ?", autobuildId).Get(&autobuild)
	if err != nil {
		sessionTK.Rollback()
		sessionMqtt.Rollback()
		return err
	}

	channelType := autobuild.AppId + p.lastfix

	// 如果mqtt状态为true，则检查是否已插入
	if autobuild.Mqtt != 0 {
		tmpMqtt := models.PushChannel{}
		_, err = sessionMqtt.Where("channel_type = ?", channelType).And("service = ?", "mqtt").And("production = ?", "storybox").Get(&tmpMqtt)
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
		ChannelType: channelType,
		Production:  "storybox",
		Service:     "mqtt",
		Params:      paramsJson,
		Status:      1,
	}
	mqtt.General.UpdatedAt = int(time.Now().Unix())
	mqtt.General.CreatedAt = int(time.Now().Unix())
	_, err = sessionMqtt.Insert(&mqtt)
	if err != nil {
		sessionTK.Rollback()
		sessionMqtt.Rollback()
		return err
	}

	autobuild.Mqtt = 1
	updateAutobuild := models.AutoBuild{
		Mqtt: 1,
	}
	_, err = sessionTK.Where("id = ?", autobuild.Id).Update(&updateAutobuild)
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
