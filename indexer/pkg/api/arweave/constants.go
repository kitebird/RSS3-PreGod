package arweave

import "time"

type ArAccount string

const (
	MirrorUploader ArAccount = "Ky1c1Kkt-jZ9sY1hvLF5nCf6WWdBhIU5Un_BMYh-t3c"
)

const (
	DefaultCrawlStep     = 10
	DefaultFromHeight    = 1
	DefaultConfirmations = 500 //TODO: Shall we make it less?
)

var DefaultCrawlConfig = &crawlConfig{
	DefaultFromHeight,
	DefaultConfirmations,
	DefaultCrawlStep,
	time.Duration(600),
}
