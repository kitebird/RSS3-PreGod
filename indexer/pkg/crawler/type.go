package crawler

import (
	"time"

	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/db/model"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/constants"
)

type Crawler interface {
	Work(WorkParam) error
	// GetResult return &{Assets, Notes, Items}
	GetResult() *CrawlerResult
	// GetBio
	GetUserBio(WorkParam) (string, error)
}

type CrawlerResult struct {
	Assets []*model.ItemId
	Notes  []*model.ItemId
	Items  []*model.Item
}

type WorkParam struct {
	Identity   string
	NetworkID  constants.NetworkID
	PlatformID constants.PlatformID // optional
	Limit      int                  // optional, aka Count, limit the number of items to be crawled

	TimeStamp   time.Time // optional
	BlockHeight int64     // optional
}

// CrawlerResult inherits the function by default

func (cr *CrawlerResult) Work(WorkParam) error {
	return nil
}

func (cr *CrawlerResult) GetResult() *CrawlerResult {
	return cr
}

func (cr *CrawlerResult) GetUserBio(WorkParam) (string, error) {
	return "", nil
}
