package item_stroge_task

import (
	"fmt"

	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/crawler"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/db"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/processor"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/rss3uri"
)

type ItemStrogeTask struct {
	processor.ProcessTaskParam
}

func NewItemStrogeParam(workParam crawler.WorkParam) *ItemStrogeTask {
	return &ItemStrogeTask{
		processor.ProcessTaskParam{
			TaskType:  processor.ProcessTaskTypeItemStroge,
			WorkParam: workParam,
		},
	}
}

func (pt *ItemStrogeTask) Fun() error {
	var err error

	var c crawler.Crawler

	var r *crawler.CrawlerResult

	instance := rss3uri.NewAccountInstance(pt.WorkParam.Identity, pt.WorkParam.PlatformID.Symbol())

	c = processor.MakeCrawlers(pt.WorkParam.NetworkID)
	if c == nil {
		err = fmt.Errorf("unsupported network id: %d", pt.WorkParam.NetworkID)

		return err
	}

	err = c.Work(pt.WorkParam)

	if err != nil {
		err = fmt.Errorf("crawler fails while working: %s", err)

		return err
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

	return nil
}
