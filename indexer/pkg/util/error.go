package util

type ErrorCode int
type ErrorMsg string

const (
	ErrorCodeSuccess             ErrorCode = iota
	ErrorCodeNotFoundData        ErrorCode = 1000
	ErrorCodeNotSupportedNetwork ErrorCode = 1001

	ErrorMsgNotFoundData        = "Not found data"
	ErrorMsgNotSupportedNetwork = "Not supported network"
)

type ErrorBase struct {
	ErrorCode ErrorCode
	ErrorMsg  ErrorMsg
}
