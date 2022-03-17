package processor

type ProcessTaskType int

const (
	ProcessTaskTypeNontType      ProcessTaskType = iota
	ProcessTaskTypeItemStroge    ProcessTaskType = 1
	ProcessTaskTypeUserBioStroge ProcessTaskType = 2
)
