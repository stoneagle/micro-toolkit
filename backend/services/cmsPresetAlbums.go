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

func (u *CmsPresetAlbums) Add(autobuildId int, albumList string) (err error) {
	sessionTK := u.engineTK.NewSession()
	defer sessionTK.Close()
	sessionAL := u.engineAL.NewSession()
	defer sessionAL.Close()

	autobuild := models.AutoBuild{}
	_, err = sessionAL.Where("id = ?", autobuildId).Get(&autobuild)
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
			AlbumId: uint(albumId),
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
