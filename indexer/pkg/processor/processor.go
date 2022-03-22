package processor

import (
	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/api/jike"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/api/moralis"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/crawler"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/constants"
)

type processTaskHandler interface {
	Fun() error
}

type ProcessTaskParam struct {
	processTaskHandler
	TaskType  ProcessTaskType
	WorkParam crawler.WorkParam
}

type ProcessTaskResult struct {
	TaskType   ProcessTaskType
	TaskResult ProcessTaskErrorCode
}

type Processor struct {
	// Emergency use, highest priority, such as user data not found
	UrgentQ chan *ProcessTaskParam
	// General use, such as access to authenticate user information
	HighQ chan *ProcessTaskParam
	// Unaffected condition use, such as polling query data
	LowQ chan *ProcessTaskParam
}

var GlobalProcessor *Processor

func Setup() error {
	GlobalProcessor = NewProcessor()
	go GlobalProcessor.ListenAndServe()

	return nil
}

func NewProcessor() *Processor {
	processor := new(Processor)

	processor.UrgentQ = make(chan *ProcessTaskParam)
	processor.HighQ = make(chan *ProcessTaskParam)
	processor.LowQ = make(chan *ProcessTaskParam)

	return processor
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

func (w *Processor) ListenAndServe() {
	select {
	case t := <-w.HighQ:
		t.Fun()
	default:
		select {
		case t := <-w.HighQ:
			t.Fun()
		case t := <-w.LowQ:
			t.Fun()
		}
	}
}
