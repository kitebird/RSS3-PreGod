package moralis

import (
	"fmt"
	"os"

	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/types"
	"github.com/joho/godotenv"
	jsoniter "github.com/json-iterator/go"
)

func GetMoralisApiKey() string {
	err := godotenv.Load(".env")
	if err != nil {
		return ""
	}
	return os.Getenv("MoralisApiKey")
}

func GetNFTs(userAddress string, chainType string, apiKey string) ([]types.NFTItem, error) {
	var headers = map[string]string{
		"accept":    "application/json",
		"X-API-Key": apiKey,
	}

	// Gets all NFT items of user
	apiUrl := fmt.Sprintf("https://deep-index.moralis.io/api/v2/%s/nft?chain=%s&format=decimal&offset=0&limit=100", userAddress, chainType)
	response, _ := Get(apiUrl, headers)

	res := new(types.NFTResult)
	jsoni := jsoniter.ConfigCompatibleWithStandardLibrary
	err := jsoni.Unmarshal(response, &res)
	if err != nil {
		return nil, nil
	}

	return res.Result, nil
}
