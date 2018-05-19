package models

type UpUpdate struct {
	GeneralWithDeleted `xorm:"extends"`
	Production         string `xorm:"varchar(200) notnull default('pudding')"`
	Local              string `xorm:"varchar(40) notnull default('zh_CN')"`
	Channel            uint   `xorm:"notnull default(10000)"`
	Name               string `xorm:"varchar(100) notnull"`
	Vcode              uint   `xorm:"notnull default(0)"`
	Vname              string `xorm:"varchar(40) notnull"`
	OSvcode            uint   `xorm:"notnull default(0)"`
	Operation          string `xorm:"varchar(1024) notnull default('')"`
	ExecuteTime        uint   `xorm:"char(50) notnull default('now')"`
	Minvcode           uint   `xorm:"notnull default(0)"`
	Maxvcode           uint   `xorm:"notnull default(0)"`
	Isforce            uint   `xorm:"smallint(3) notnull default(0)"`
	Filename           string `xorm:"varchar(100) notnull"`
	Filesize           uint   `xorm:"notnull"`
	Dwnpath            string `xorm:"varchar(1024) notnull"`
	DwnpathBk          string `xorm:"varchar(1024) notnull"`
	Md5sum             uint   `xorm:"char(32) notnull"`
	Feature            string `xorm:"varchar(1024) default(NULL)"`
	Ownname            string `xorm:"varchar(100) default(NULL)"`
}

func (c UpUpdate) TableName() string {
	return "up_update"
}
