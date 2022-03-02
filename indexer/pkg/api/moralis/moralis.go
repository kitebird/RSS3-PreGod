package moralis

import (
	"fmt"
	"strings"

	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/util"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/types"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/config"
	jsoniter "github.com/json-iterator/go"
)

type (
	MoralisNFTResult         = types.MoralisNFTResult
	MoralisNFTTransferResult = types.MoralisNFTTransferResult
	MoralisGetLogsResult     = types.MoralisGetLogsResult
)

const endpoint = "https://deep-index.moralis.io"

var jsoni = jsoniter.ConfigCompatibleWithStandardLibrary

func GetMoralisApiKey() string {
	if err := config.Setup(); err != nil {
		return ""
	}

	apiKey, err := jsoni.MarshalToString(config.Config.Indexer.Moralis.ApiKey)
	if err != nil {
		return ""
	}

	return strings.Trim(apiKey, "\"")
}

func GetNFTs(userAddress string, chainType string, apiKey string) (MoralisNFTResult, error) {
	var headers = map[string]string{
		"accept":    "application/json",
		"X-API-Key": apiKey,
	}

	// Gets all NFT items of user
	url := fmt.Sprintf("%s/api/v2/%s/nft?chain=%s&format=decimal",
		endpoint, userAddress, chainType)

	response, err := util.Get(url, headers)
	if err != nil {
		return MoralisNFTResult{}, err
	}

	res := new(MoralisNFTResult)

	err = jsoni.Unmarshal(response, &res)
	if err != nil {
		return MoralisNFTResult{}, err
	}

	return *res, nil
}

func GetNFTTransfers(userAddress string, chainType string, apiKey string) (MoralisNFTTransferResult, error) {
	var headers = map[string]string{
		"accept":    "application/json",
		"X-API-Key": apiKey,
	}

	// Gets all NFT transfers of user
	url := fmt.Sprintf("%s/api/v2/%s/nft/transfers?chain=%s&format=decimal&direction=both",
		endpoint, userAddress, chainType)

	response, err := util.Get(url, headers)
	if err != nil {
		return MoralisNFTTransferResult{}, err
	}

	res := new(MoralisNFTTransferResult)

	err = jsoni.Unmarshal(response, &res)
	if err != nil {
		return MoralisNFTTransferResult{}, err
	}

	return *res, nil
}

func GetLogs(fromBlock int64, toBlock int64, address string, topic string, chainType string, apiKey string) (MoralisGetLogsResult, error) {
	var headers = map[string]string{
		"accept":    "application/json",
		"X-API-Key": apiKey,
	}

	url := fmt.Sprintf("%s/api/v2/%s/logs?chain=%s&from_block=%d&to_block=%d&topic0=%s",
		endpoint, address, chainType, fromBlock, toBlock, topic)
	response, err := util.Get(url, headers)
	if err != nil {
		return MoralisGetLogsResult{}, err
	}
	//fmt.Println(string(response))

	res := new(MoralisGetLogsResult)

	err = jsoni.Unmarshal(response, &res)
	if err != nil {
		return MoralisGetLogsResult{}, err
	}

	return *res, nil
}
