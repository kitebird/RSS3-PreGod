package worker

import (
	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/crawlers"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/db"
)

func ProcessTask(t *Task) error {
	mw := crawlers.NewMoralisCrawler()
	err := mw.Work(t.Identity, t.Network)
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
	//TODO: save by account: <identity>@<platform>
	db.SetAssets(t.Identity, assets)
	db.AppendNotes(t.Identity, notes)

	return nil
}
