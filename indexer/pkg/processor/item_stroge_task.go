package processor

import (
	"fmt"

	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/crawler"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/db"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/logger"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/rss3uri"
)

type itemStrogeTask struct {
	ProcessTaskParam
}

func NewItemStrogeParam() ProcessTaskUnit {
	return &itemStrogeTask{
		ProcessTaskParam{
			TaskType: ProcessTaskTypeItemStroge,
		},
	}
}

func (pt *itemStrogeTask) Fun() error {
	var err error

	var c crawler.Crawler

	var r *crawler.CrawlerResult

	instance := rss3uri.NewAccountInstance(pt.WorkParam.Identity, pt.WorkParam.PlatformID.Symbol())

	c = MakeCrawlers(pt.WorkParam.NetworkID)
	if c == nil {
		err = fmt.Errorf("unsupported network id: %d", pt.WorkParam.NetworkID)

		goto RETURN
	}

	err = c.Work(pt.WorkParam)

	if err != nil {
		err = fmt.Errorf("crawler fails while working: %s", err)

		goto RETURN
	}

	r = c.GetResult()
	if r.Items != nil {
		for _, item := range r.Items {
			db.InsertItem(item)
		}
	}

	if r.Assets != nil {
		db.SetAssets(instance, r.Assets, pt.WorkParam.NetworkID)
	}

	if r.Notes != nil {
		db.AppendNotes(instance, r.Notes)
	}

RETURN:
	if err != nil {
		logger.Error(err)

		return err
	} else {
		return nil
	}
}
