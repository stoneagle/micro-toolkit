package models

type UpUpdate struct {
	General     `xorm:"extends"`
	DeletedAt   int    `xorm:"INT(11) deleted comment('软删除时间')"`
	Production  string `xorm:"varchar(200) notnull default('pudding')"`
	Local       string `xorm:"varchar(40) notnull default('zh_CN')"`
	Channel     int    `xorm:"notnull default(10000)"`
	Name        string `xorm:"varchar(100) notnull"`
	Vcode       int    `xorm:"notnull default(0)"`
	Vname       string `xorm:"varchar(40) notnull"`
	Osvcode     int    `xorm:"notnull default(0)"`
	Operation   string `xorm:"varchar(1024) notnull default('')"`
	ExecuteTime int    `xorm:"char(50) notnull default('now')"`
	Minvcode    int    `xorm:"notnull default(0)"`
	Maxvcode    int    `xorm:"notnull default(0)"`
	Isforce     int    `xorm:"smallint(3) notnull default(0)"`
	Filename    string `xorm:"varchar(100) notnull"`
	Filesize    int    `xorm:"notnull"`
	Dwnpath     string `xorm:"varchar(1024) notnull"`
	DwnpathBk   string `xorm:"varchar(1024) notnull"`
	Md5sum      int    `xorm:"char(32) notnull"`
	Feature     string `xorm:"varchar(1024) default(NULL)"`
	Ownname     string `xorm:"varchar(100) default(NULL)"`
}

func (c UpUpdate) TableName() string {
	return "up_update"
}
