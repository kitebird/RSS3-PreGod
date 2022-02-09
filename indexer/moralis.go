package indexer

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"

	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/types"
)

var API_KEY = "" // maybe read api key from config file
var HEADERS = map[string]string{
	"accept":    "application/json",
	"X-API-Key": API_KEY,
}
var jsoni = jsoniter.ConfigCompatibleWithStandardLibrary

func GetNFTs(userAddress string, chainType string) ([]types.NFTItem, error) {
	// Gets all NFT items of user
	// ETH
	api_url := fmt.Sprintf("https://deep-index.moralis.io/api/v2/%s/nft?chain=%s&format=decimal&offset=0&limit=100", userAddress, chainType)
	response, _ := Get(api_url, HEADERS)
	//fmt.Println(res)

	res := new(types.NFTResult)
	err := jsoni.Unmarshal(response, &res)
	if err != nil {
		return nil, nil
	}

	return res.Result, nil
}
