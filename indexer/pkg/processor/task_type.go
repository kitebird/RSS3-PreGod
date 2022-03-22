package processor

type ProcessTaskType int

const (
	ProcessTaskTypeNontType      ProcessTaskType = iota
	ProcessTaskTypeItemStroge    ProcessTaskType = 1
	ProcessTaskTypeUserBioStroge ProcessTaskType = 2
)

type ProcessTaskErrorCode int

const (
	ProcessTaskErrorCodeSuccess   ProcessTaskErrorCode = iota
	ProcessTaskErrorCodeFoundData ProcessTaskErrorCode = 1
)
