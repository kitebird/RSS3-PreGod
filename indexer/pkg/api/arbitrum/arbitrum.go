package arbitrum

import (
	"fmt"
	"strings"

	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/util"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/config"
	jsoniter "github.com/json-iterator/go"
	"github.com/valyala/fastjson"
)

const endpoint = "https://api.arbiscan.io"

var jsoni = jsoniter.ConfigCompatibleWithStandardLibrary

func GetApiKey() string {
	if err := config.Setup(); err != nil {
		return ""
	}

	apiKey, err := jsoni.MarshalToString(config.Config.Indexer.Aribtrum.ApiKey)
	if err != nil {
		return ""
	}

	return strings.Trim(apiKey, "\"")
}

func GetNFTTransfers(owner string) ([]NFTTransferItem, error) {
	apiKey := GetApiKey()
	url := fmt.Sprintf("%s/api?module=account&action=tokennfttx&address=%s&startblock=0&endblock=999999999&sort=asc&apikey=%s", endpoint, owner, apiKey)

	response, err := util.Get(url, nil)
	if err != nil {
		return nil, nil
	}

	var parser fastjson.Parser

	parsedJson, parseErr := parser.Parse(string(response))
	if parseErr != nil {
		return nil, nil
	}

	result := make([]NFTTransferItem, 0)

	arrys := parsedJson.GetArray("result")
	for _, v := range arrys {
		var item NFTTransferItem
		item.Address = string(v.GetStringBytes("contractAddress"))
		item.TokenId = string(v.GetStringBytes("tokenID"))
		item.Name = string(v.GetStringBytes("tokenName"))
		item.Symbol = string(v.GetStringBytes("tokenSymbol"))
		item.From = string(v.GetStringBytes("from"))
		item.To = string(v.GetStringBytes("to"))
		item.TimeStamp = string(v.GetStringBytes("timeStamp"))
		item.Hash = string(v.GetStringBytes("hash"))

		result = append(result, item)
	}

	return result, nil
}
