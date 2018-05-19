package models

type AutoBuild struct {
	GeneralWithDeleted `xorm:"extends"`
	AppId              string `xorm:"varchar(32) notnull unique comment('目标App')"`
	CmsSourceApp       string `xorm:"varchar(32) default('') unique comment('cms来源App')"`
	CmsExecTime        string `xorm:"varchar(32) default('') unique comment('cms执行时间')"`
	Mqtt               bool   `xorm:"default(0) comment('Mqtt状态:0未配置1已配置')"`
	Callback           uint   `xorm:"default(0) comment('callback状态:0未配置1rtoy2童秘')"`
	UpgradeName        string `xorm:"default('') comment('升级配置名称')"`
	UpgradeVcode       uint   `xorm:"default(0) comment('升级配置号')"`
	UpgradeVname       string `xorm:"default('') comment('升级配置展示名称')"`
	AlbumList          string `xorm:"default(0) comment('预置歌单列表,分隔')"`
}

func (c AutoBuild) TableName() string {
	return "auto_build"
}
