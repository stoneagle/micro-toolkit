package common

type ErrorCode int

const (
	ErrorOk     ErrorCode = 0
	ErrorParams ErrorCode = 1
	ErrorMysql  ErrorCode = 2
	ErrorLogin  ErrorCode = 3
	ErrorServer ErrorCode = 500
)
