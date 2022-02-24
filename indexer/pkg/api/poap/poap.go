package poap

import (
	"fmt"

	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/util"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/types"
	jsoniter "github.com/json-iterator/go"
)

const endpoint = "https://api.poap.xyz"

var jsoni = jsoniter.ConfigCompatibleWithStandardLibrary

func GetActions(user string) (types.PoapResponse, error) {
	url := fmt.Sprintf("%s/actions/scan/%s",
		endpoint, user)
	response, _ := util.Get(url, nil)

	res := new(types.PoapResponse)

	err := jsoni.Unmarshal(response, &res)
	if err != nil {
		return types.PoapResponse{}, err
	}

	return *res, nil
}

func GetAction(tokenId string) (types.PoapEventInfo, error) {

	url := fmt.Sprintf("%s/token/%s",
		endpoint, tokenId)
	response, _ := util.Get(url, nil)

	res := new(types.TokenResponse)

	err := jsoni.Unmarshal(types.TokenResponse.PoapEventInfo, &res)
	if err != nil {
		return types.PoapEventInfo{}, err
	}

	// res.PoapEventInfo.StartDate =
	// res.PoapEventInfo.EndDate =
	// res.PoapEventInfo.ExpiryDate =

	return *res, nil

}
