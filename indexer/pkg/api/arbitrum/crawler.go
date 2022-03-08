package arbitrum

import (
	"time"

	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/crawler"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/db/model"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/constants"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/logger"
)

type abCrawler struct {
	crawler.CrawlerResult
}

func NewArbitrumCrawler() crawler.Crawler {
	return &abCrawler{
		crawler.CrawlerResult{
			Assets:  []*model.ItemId{},
			Notes:   []*model.ItemId{},
			Items:   []*model.Item{},
			Objects: []*model.Object{},
		},
	}
}
func (ac *abCrawler) Work(userAddress string, itemType constants.NetworkName) error {
	nftTransfers, err := GetNFTTransfers(userAddress)
	if err != nil {
		return err
	}

	assets, err := GetNFTs(userAddress)
	if err != nil {
		return err
	}

	// parse nft transfers
	for _, v := range nftTransfers {
		tsp, err := time.Parse(time.RFC3339, v.TimeStamp)
		if err != nil {
			tsp = time.Now()
		}

		ni := model.NewItem(
			v.GetUid(),
			constants.ItemType_Arbitrum_Nft,
			v.From,
			v.To,
			v.Hash,
			tsp,
		)
		ac.Items = append(ac.Items, ni)
		ac.Notes = append(ac.Notes, &model.ItemId{
			ItemTypeID: constants.ItemType_Arbitrum_Nft,
			Proof:      v.Hash,
		})
	}

	// parse nfts
	for _, v := range assets {
		obj := model.NewObject(nil, v.GetUid(), constants.ItemType_Arbitrum_Nft, "", "", nil, nil)
		ac.Objects = append(ac.Objects, obj)

		hasProof := false

		for _, nftTransfer := range nftTransfers {
			if nftTransfer.EqualsToToken(v) {
				hasProof = true

				ac.Assets = append(ac.Assets, &model.ItemId{
					ItemTypeID: constants.ItemType_Arbitrum_Nft,
					Proof:      nftTransfer.Hash,
				})
			}
		}

		if !hasProof {
			logger.Error("Asset doesn't has proof")
		}
	}

	return nil
}

func (ac *abCrawler) GetResult() *crawler.CrawlerResult {
	return &crawler.CrawlerResult{
		Assets:  ac.Assets,
		Notes:   ac.Notes,
		Items:   ac.Items,
		Objects: ac.Objects,
	}
}
