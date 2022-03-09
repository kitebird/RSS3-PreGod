package processor

import (
	"fmt"

	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/api/jike"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/api/moralis"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/crawler"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/db"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/constants"
)

type worker struct {
	lowQ, highQ chan *Task
}

type Worker interface {
	processTask(t *Task) error

	ListenAndServe()
}

func NewWorker(lq, hq chan *Task) Worker {
	return &worker{lq, hq}
}

func makeCrawlers(network constants.NetworkID) crawler.Crawler {
	switch network {
	case constants.NetworkIDEthereumMainnet,
		constants.NetworkIDBNBChain,
		constants.NetworkIDAvalanche,
		constants.NetworkIDFantom,
		constants.NetworkIDPolygon:
		return moralis.NewMoralisCrawler()
	case constants.NetworkIDJike:
		return jike.NewJikeCrawler()
	default:
		return nil
	}
}

func (w *worker) processTask(t *Task) error {
	c := makeCrawlers(t.Network)
	if c == nil {
		return fmt.Errorf("unsupported network: %d", t.Network)
	}

	err := c.Work(t.Identity, t.Network)

	if err != nil {
		panic(err)
	}

	r := c.GetResult()

	for _, item := range r.Items {
		db.InsertItemDoc(item)
	}
	//TODO: save by account: <identity>@<platform>
	db.SetAssets(t.Identity, r.Assets)
	db.AppendNotes(t.Identity, r.Notes)

	return nil
}

func (w *worker) ListenAndServe() {
	select {
	case t := <-w.highQ:
		w.processTask(t)
	default:
		select {
		case t := <-w.highQ:
			w.processTask(t)
		case t := <-w.lowQ:
			w.processTask(t)
		}
	}
}
