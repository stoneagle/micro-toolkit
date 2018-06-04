package models

type General struct {
	Id        int `xorm:"pk autoincr"`
	CreatedAt int `INT(11) xorm:"created comment('创建时间')"`
	UpdatedAt int `INT(11) xorm:"updated comment('修改时间')"`
}
