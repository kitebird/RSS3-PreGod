package crawlers

import (
	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/db/model"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/constants"
)

type Crawler interface {
	Work(string, constants.NetworkName) error
	// GetResult return assets, notes, items, objects
	GetResult() ([]*model.ItemId, []*model.ItemId, []*model.Item, []*model.Object)

	GetAssets() []*model.ItemId
	GetNotes() []*model.ItemId
	GetItems() []*model.Item
	GetObjects() []*model.Object
}
