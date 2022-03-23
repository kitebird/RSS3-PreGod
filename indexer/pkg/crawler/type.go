package crawler

import (
	"time"

	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/db/model"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/constants"
)

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

	LastIndexedTsp time.Time // optional, if provided, only index items newer than this time
}

type Crawler interface {
	Work(WorkParam) (*CrawlerResult, error)
}
