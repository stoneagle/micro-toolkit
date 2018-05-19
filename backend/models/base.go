package models

import "time"

type GeneralWithDeleted struct {
	Id        uint      `xorm:"pk autoincr"`
	CreatedAt time.Time `xorm:"created comment('创建时间')"`
	UpdatedAt time.Time `xorm:"updated comment('修改时间')"`
	DeletedAt time.Time `xorm:"deleted comment('软删除时间')"`
}

type General struct {
	Id        uint      `xorm:"pk autoincr"`
	CreatedAt time.Time `xorm:"created comment('创建时间')"`
	UpdatedAt time.Time `xorm:"updated comment('修改时间')"`
}
