package worker

import (
	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/crawlers/moralis"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/db"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/constants"
)

func ProcessResult(userAddress string, itemType constants.ItemTypeID) error {
	mw := moralis.NewMoralisCrawler()
	err := mw.Work(userAddress, itemType)
	if err != nil {
		panic(err)
	}
	assets, notes, items, objects := mw.GetResult()

	for _, item := range items {
		db.InsertItemDoc(item)
	}
	for _, object := range objects {
		db.InsertObjectDoc(object)
	}
	db.SetAssets(userAddress, assets)
	db.AppendNotes(userAddress, notes)

	return nil
}
