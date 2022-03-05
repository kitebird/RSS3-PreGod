package poap

import (
	"fmt"
	"time"

	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/crawler"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/db/model"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/constants"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/logger"
)

type poapCralwer struct {
	rss3Items   []*model.Item
	rss3Objects []*model.Object

	rss3Assets, rss3Notes []*model.ItemId
}

func NewPoapCralwer() crawler.Crawler {
	return &poapCralwer{
		rss3Items:   []*model.Item{},
		rss3Objects: []*model.Object{},

		rss3Assets: []*model.ItemId{},
		rss3Notes:  []*model.ItemId{},
	}
}

type ChainType string

const (
	Gnosis ChainType = "Gnosis"
)

func (pc *poapCralwer) Work(userAddress string, itemType constants.NetworkName) error {
	if itemType != constants.NetworkName_Gnosis {
		return fmt.Errorf("network is not gnosis")
	}

	poapResps, err := GetActions(userAddress)
	if err != nil {
		return err
	}

	//TODO: Since we are getting the full amount of interfaces,
	// I hope to get incremental interfaces in the future and use other methods to improve efficiency
	for _, poapResp := range poapResps {
		tsp, err := poapResp.GetTsp()
		if err != nil {
			// TODO: log error
			logger.Error(tsp, err)
			tsp = time.Now()
		}

		ni := model.NewItem(
			poapResp.TokenId,
			constants.ItemType_Xdai_Poap,
			"0x0", // temp
			poapResp.Owner,
			"0x0",
			tsp,
		)
		pc.rss3Items = append(pc.rss3Items, ni)
		pc.rss3Notes = append(pc.rss3Notes, &model.ItemId{
			ItemTypeID: constants.ItemType_Xdai_Poap,
			Proof:      "Here is the proof",
		})

		pc.rss3Assets = append(pc.rss3Notes, &model.ItemId{
			ItemTypeID: constants.ItemType_Xdai_Poap,
			Proof:      "Here is the proof",
		})
	}

	return nil
}

func (pc *poapCralwer) GetResult() *crawler.CrawlerResult {
	return &crawler.CrawlerResult{
		Assets:  pc.rss3Assets,
		Notes:   pc.rss3Notes,
		Items:   pc.rss3Items,
		Objects: pc.rss3Objects,
	}
}
