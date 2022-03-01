package poap

import (
	"fmt"

	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/util"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/types"
	jsoniter "github.com/json-iterator/go"
)

const endpoint = "https://api.poap.xyz"

var jsoni = jsoniter.ConfigCompatibleWithStandardLibrary

func GetActions(user string) ([]types.PoapResponse, error) {
	if len(user) == 0 {
		var err = fmt.Errorf("user address is empty")

		return []types.PoapResponse{}, err
	}

	url := fmt.Sprintf("%s/actions/scan/%s",
		endpoint, user)
	response, err := util.Get(url, nil)
	if err != nil {
		return []types.PoapResponse{}, err
	}

	res := new([]types.PoapResponse)

	err = jsoni.Unmarshal(response, &res)
	if err != nil {
		return []types.PoapResponse{}, err
	}

	return *res, nil
}
