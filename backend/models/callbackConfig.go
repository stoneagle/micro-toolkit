package models

type CallbackConfig struct {
	General       `xorm:"extends"`
	AppId         string `xorm:"varchar(50) notnull"`
	State         int    `xorm:"notnull default(0)"`
	CallbackUrl   string `xorm:"varchar(300) notnull"`
	CallbackState string `xorm:"varchar(50) notnull"`
	Action        string `xorm:"varchar(50) notnull default('')"`
}

type CallbackTemplate struct {
	Url    string
	State  string
	Action string
}

func (c CallbackConfig) TableName() string {
	return "callback_config"
}
