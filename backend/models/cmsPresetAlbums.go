package models

type CmsPresetAlbums struct {
	GeneralWithDeleted `xorm:"extends"`
	AppId              string `xorm:"varchar(255) notnull default('') 'appId'"`
	AlbumId            int    `xorm:"BIGINT(20) notnull default(0)"`
}

func (c CmsPresetAlbums) TableName() string {
	return "cms_preset_albums"
}
