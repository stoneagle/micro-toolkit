package services

import (
	"strconv"
	"strings"
	"toolkit/backend/models"

	"github.com/go-xorm/xorm"
)

type CmsPresetAlbums struct {
	engineTK *xorm.Engine
	engineAL *xorm.Engine
}

func NewCmsPresetAlbums(tk, al *xorm.Engine) *CmsPresetAlbums {
	return &CmsPresetAlbums{
		engineTK: tk,
		engineAL: al,
	}
}

func (u *CmsPresetAlbums) List(appId string) (albums []models.CmsPresetAlbums, err error) {
	albums = make([]models.CmsPresetAlbums, 0)
	err = u.engineAL.Where("appId = ?", appId).Find(&albums)
	return
}

func (s *CmsPresetAlbums) Delete(appId string) (err error) {
	session := s.engineAL.NewSession()
	defer session.Close()
	lists, err := s.List(appId)
	if err != nil {
		session.Rollback()
		return
	}
	for _, one := range lists {
		_, err = session.Id(one.GeneralWithDeleted.Id).Delete(&one)
		if err != nil {
			session.Rollback()
			return
		}
	}
	err = session.Commit()
	return
}

func (u *CmsPresetAlbums) Add(autobuildId int, albumList string) (err error) {
	sessionTK := u.engineTK.NewSession()
	defer sessionTK.Close()
	sessionAL := u.engineAL.NewSession()
	defer sessionAL.Close()

	autobuild := models.AutoBuild{}
	_, err = sessionTK.Where("id = ?", autobuildId).Get(&autobuild)
	if err != nil {
		sessionTK.Rollback()
		sessionAL.Rollback()
		return err
	}

	albumStrSlice := strings.Split(albumList, ",")
	for _, albumIdStr := range albumStrSlice {
		albumId, err := strconv.Atoi(albumIdStr)
		if err != nil {
			return err
		}
		album := models.CmsPresetAlbums{
			AppId:   autobuild.AppId,
			AlbumId: albumId,
		}
		_, err = sessionAL.Insert(&album)
		if err != nil {
			sessionTK.Rollback()
			sessionAL.Rollback()
			return err
		}
	}

	// TODO,增量配置
	updateAutobuild := models.AutoBuild{
		AlbumList: albumList,
	}
	_, err = sessionTK.Where("id = ?", autobuild.Id).Update(&updateAutobuild)
	if err != nil {
		sessionTK.Rollback()
		sessionAL.Rollback()
		return err
	}

	err = sessionTK.Commit()
	if err != nil {
		sessionTK.Rollback()
		sessionAL.Rollback()
		return err
	}
	err = sessionAL.Commit()
	if err != nil {
		sessionTK.Rollback()
		sessionAL.Rollback()
		return err
	}
	return nil
}
