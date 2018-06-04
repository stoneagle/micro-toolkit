package models

type GeneralWithDeleted struct {
	Id        int `xorm:"pk autoincr"`
	CreatedAt int `xorm:"INT(11) created comment('创建时间')"`
	UpdatedAt int `xorm:"INT(11) updated comment('修改时间')"`
	DeletedAt int `xorm:"INT(11) deleted comment('软删除时间')"`
}

type General struct {
	Id        int `xorm:"pk autoincr"`
	CreatedAt int `INT(11) xorm:"created comment('创建时间')"`
	UpdatedAt int `INT(11) xorm:"updated comment('修改时间')"`
}
