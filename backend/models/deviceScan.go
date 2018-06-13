package models

type DeviceScan struct {
	General  `xorm:"extends"`
	AppId    string `xorm:"varchar(255) notnull default('')"`
	ScanType int    `xorm:"BIGINT(11) notnull default(0)"`
	FuncType int    `xorm:"BIGINT(11) notnull default(0)"`
	InfoType int    `xorm:"BIGINT(11) notnull default(0)"`
	Info     string `xorm:"varchar(512) notnull default('')"`
}
