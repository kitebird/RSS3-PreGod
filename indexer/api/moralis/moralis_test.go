package moralis

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_get(t *testing.T) {
	apiKey := GetMoralisApiKey()
	items, err := GetNFTs("0x3b6d02a24df681ffdf621d35d70aba7adaac07c1", "eth", apiKey)
	for _, item := range items {
		fmt.Println(item)
	}
	// assert for nil
	assert.Nil(t, err)
	// assert equality
	//assert.Equal(t, len(items), 5, "they should be equal")
}
