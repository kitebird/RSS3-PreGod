package jike

import (
	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/crawler"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/db/model"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/constants"
)

type jikeCrawler struct {
	rss3Items   []*model.Item
	rss3Objects []*model.Object

	rss3Assets, rss3Notes []*model.ItemId
}

func NewJikeCrawler() crawler.Crawler {
	return &jikeCrawler{
		rss3Items:   []*model.Item{},
		rss3Objects: []*model.Object{},
		rss3Notes:   []*model.ItemId{},
	}
}

func (mc *jikeCrawler) Work(userAddress string, itemType constants.NetworkName) error {
	timeline, err := GetUserTimeline(userAddress)

	if err != nil {
		return err
	}

	for _, item := range timeline {
		ni := model.NewItem(
			item.Link, // object uid for jike
			constants.ItemType_Jike_Node,
			item.Author,
			"",
			item.Link,
			item.Timestamp,
		)
		mc.rss3Items = append(mc.rss3Items, ni)

		mc.rss3Notes = append(mc.rss3Notes, &model.ItemId{
			ItemTypeID: constants.ItemType_Jike_Node,
			Proof:      item.Link,
		})

		no := model.NewObject(
			[]string{item.Author},
			item.Link, // object uid for jike
			constants.ItemType_Jike_Node,
			item.Summary, // use title as summary
			item.Summary,
			[]string{"Jike Post"},
			item.Attachments,
		)
		mc.rss3Objects = append(mc.rss3Objects, no)
	}

	return nil
}

func (mc *jikeCrawler) GetResult() *crawler.CrawlerResult {
	return &crawler.CrawlerResult{
		Assets:  mc.rss3Assets,
		Notes:   mc.rss3Notes,
		Items:   mc.rss3Items,
		Objects: mc.rss3Objects,
	}
}
