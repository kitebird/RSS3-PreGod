package crawlers

import "github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/db/model"

type poapCralwer struct {
	rss3Items   []*model.Item
	rss3Objects []*model.Object

	rss3Assets, rss3Notes []*model.ItemId
}

func NewPoapCralwer() {
	return &poapCralwer{
		rss3Items:   []*model.Item{},
		rss3Objects: []*model.Object{},

		rss3Assets: []*model.ItemId{},
		rss3Notes:  []*model.ItemId{},
	}
}
