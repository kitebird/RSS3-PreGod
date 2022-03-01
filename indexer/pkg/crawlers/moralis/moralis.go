package moralis

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/crawlers"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/db/model"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/util"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/types"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/constants"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/logger"
	"github.com/joho/godotenv"
	jsoniter "github.com/json-iterator/go"
)

const endpoint = "https://deep-index.moralis.io"

type moralisCralwer struct {
	rss3Items   []*model.Item
	rss3Objects []*model.Object

	rss3Assets, rss3Notes []*model.ItemId
}

func NewMoralisCrawler() crawlers.Crawler {
	return &moralisCralwer{
		rss3Items:   []*model.Item{},
		rss3Objects: []*model.Object{},

		rss3Assets: []*model.ItemId{},
		rss3Notes:  []*model.ItemId{},
	}
}

func (mc *moralisCralwer) Work(userAddress string, itemType constants.ItemTypeID) error {
	nftTransfers, err := GetNFTTransfers(userAddress, ETH, GetMoralisApiKey())
	if err != nil {
		return err
	}
	log.Println(nftTransfers.Total)
	//TODO: tsp

	assets, err := GetNFTs(userAddress, ETH, GetMoralisApiKey())
	if err != nil {
		return err
	}
	//parser
	for _, nftTransfer := range nftTransfers.Result {
		tsp, err := nftTransfer.GetTsp()
		if err != nil {
			// TODO: log error
			logger.Error(tsp, err)
			tsp = time.Now()
		}
		ni := model.NewItem(
			nftTransfer.GetUid(),
			itemType,
			nftTransfer.FromAddress,
			nftTransfer.ToAddress,
			nftTransfer.TransactionHash,
			tsp,
		)
		mc.rss3Items = append(mc.rss3Items, ni)
		mc.rss3Notes = append(mc.rss3Notes, &model.ItemId{
			ItemType: itemType,
			Proof:    nftTransfer.TransactionHash,
		})
	}
	for _, asset := range assets.Result {
		// TODO: make attachments and authors
		no := model.NewObject(nil, asset.GetUid(), itemType, "", "", nil, nil)
		mc.rss3Objects = append(mc.rss3Objects, no)
		hasProof := false
		for _, nftTransfer := range nftTransfers.Result {
			if nftTransfer.EqualsToToken(asset) {
				hasProof = true
				mc.rss3Assets = append(mc.rss3Assets, &model.ItemId{
					ItemType: itemType,
					Proof:    nftTransfer.TransactionHash,
				})
			}
		}
		if !hasProof {
			// TODO: error handle here
			logger.Errorf("Asset doesn't has proof.")
		}
	}
	// make the object list complete
	for _, nftTransfer := range nftTransfers.Result {
		hasObject := false
		for _, asset := range assets.Result {
			if nftTransfer.EqualsToToken(asset) && asset.MetaData != "" {
				hasObject = true
			}
		}
		if !hasObject {
			// TODO: get object
			logger.Errorf("Asset doesn't has the metadata.")
		}
	}

	return nil
}

func (mc *moralisCralwer) GetResult() ([]*model.ItemId, []*model.ItemId, []*model.Item, []*model.Object) {
	return mc.rss3Assets, mc.rss3Notes, mc.rss3Items, mc.rss3Objects
}

func (mc *moralisCralwer) GetAssets() []*model.ItemId {
	return mc.rss3Assets
}

func (mc *moralisCralwer) GetNotes() []*model.ItemId {
	return mc.rss3Notes
}

func (mc *moralisCralwer) GetItems() []*model.Item {
	return mc.rss3Items
}

func (mc *moralisCralwer) GetObjects() []*model.Object {
	return mc.rss3Objects
}

var jsoni = jsoniter.ConfigCompatibleWithStandardLibrary

func GetMoralisApiKey() string {
	if err := godotenv.Load(".env"); err != nil {
		return ""
	}

	return os.Getenv("MoralisApiKey")
}

func GetNFTs(userAddress string, chainType MoralisChainType, apiKey string) (types.MoralisNFTResult, error) {
	var headers = map[string]string{
		"accept":    "application/json",
		"X-API-Key": apiKey,
	}

	// Gets all NFT items of user
	url := fmt.Sprintf("%s/api/v2/%s/nft?chain=%s&format=decimal",
		endpoint, userAddress, chainType)
	response, _ := util.GetURL(url, headers)

	res := new(types.MoralisNFTResult)

	err := jsoni.Unmarshal(response, &res)
	if err != nil {
		return types.MoralisNFTResult{}, err
	}

	return *res, nil
}

func GetNFTTransfers(userAddress string, chainType MoralisChainType, apiKey string) (types.MoralisNFTTransferResult, error) {
	var headers = map[string]string{
		"accept":    "application/json",
		"X-API-Key": apiKey,
	}

	// Get all NFT transfers of user
	url := fmt.Sprintf("%s/api/v2/%s/nft/transfers?chain=%s&format=decimal&direction=both",
		endpoint, userAddress, chainType)
	response, _ := util.GetURL(url, headers)

	res := new(types.MoralisNFTTransferResult)

	err := jsoni.Unmarshal(response, &res)
	if err != nil {
		return types.MoralisNFTTransferResult{}, err
	}

	return *res, nil
}
