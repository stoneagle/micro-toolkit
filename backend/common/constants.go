package common

type ErrorCode int

const (
	ErrorOk     ErrorCode = 0
	ErrorParams ErrorCode = 1
	ErrorServer ErrorCode = 500
)
