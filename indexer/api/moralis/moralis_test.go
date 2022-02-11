package moralis

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetNFT(t *testing.T) {
	apiKey := GetMoralisApiKey()
	result, err := GetNFTs("0x3b6d02a24df681ffdf621d35d70aba7adaac07c1", "eth", apiKey)

	// assert for nil
	assert.Nil(t, err)

	for _, item := range result.Result {
		fmt.Println(item)
	}

	// assert equality
	//assert.Equal(t, len(items), 5, "they should be equal")
}

func Test_GetNFTTransfers(t *testing.T) {
	apiKey := GetMoralisApiKey()
	result, err := GetNFTTransfers("0x3b6d02a24df681ffdf621d35d70aba7adaac07c1", "eth", apiKey)

	// assert for nil
	assert.Nil(t, err)

	for _, item := range result.Result {
		fmt.Println(item)
	}

	// assert equality
	//assert.Equal(t, len(items), 5, "they should be equal")
}
