package processor

import (
	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/api/jike"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/api/moralis"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/crawler"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/constants"
)

type ProcessTaskParam struct {
	TaskType  ProcessTaskType
	WorkParam crawler.WorkParam
}

type ProcessTaskUnit interface {
	Fun() error
}

type Processor interface {
	ListenAndServe()
}

type processor struct {
	// Emergency use, highest priority, such as user data not found
	UrgentQ chan ProcessTaskUnit
	// General use, such as access to authenticate user information
	highQ chan ProcessTaskUnit
	// Unaffected condition use, such as polling query data
	lowQ chan ProcessTaskUnit
}

func NewProcessor(uq, lq, hq chan ProcessTaskUnit) Processor {
	return &processor{uq, lq, hq}
}

func NewTaskQueue() chan ProcessTaskUnit {
	return make(chan ProcessTaskUnit)
}

func MakeCrawlers(network constants.NetworkID) crawler.Crawler {
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

func (w *processor) ListenAndServe() {
	select {
	case t := <-w.highQ:
		t.Fun()
	default:
		select {
		case t := <-w.highQ:
			t.Fun()
		case t := <-w.lowQ:
			t.Fun()
		}
	}
}
