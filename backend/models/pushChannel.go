package models

type PushChannel struct {
	General     `xorm:"extends"`
	ChannelType string `xorm:"varchar(64) notnull default(0)"`
	Production  string `xorm:"varchar(64) notnull"`
	Service     string `xorm:"varchar(256) notnull default('')"`
	Ciduniq     uint   `xorm:"smallint(4) default(0)"`
	Uniqkey     uint   `xorm:"char(8) default(0)"`
	Prority     uint   `xorm:"smallint(4) notnull default(0)"`
	Params      string `xorm:"TEXT"`
	Status      uint   `xorm:"notnull default(1)"`
	Desc        string `xorm:"varchar(256) notnull"`
}

func (c PushChannel) TableName() string {
	return "push_channels"
}
