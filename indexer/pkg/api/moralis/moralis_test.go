package moralis_test

import (
	"testing"

	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/api/moralis"
	"github.com/stretchr/testify/assert"
)

func TestGetNFT(t *testing.T) {
	t.Parallel()

	apiKey := moralis.GetMoralisApiKey()
	result, err := moralis.GetNFTs("0x3b6d02a24df681ffdf621d35d70aba7adaac07c1", "eth", apiKey)
	// assert for nil
	assert.Nil(t, err)

	assert.True(t, len(result.Result) > 0)
}

func TestGetNFTTransfers(t *testing.T) {
	t.Parallel()

	apiKey := moralis.GetMoralisApiKey()
	result, err := moralis.GetNFTTransfers("0x3b6d02a24df681ffdf621d35d70aba7adaac07c1", "eth", apiKey)
	// assert for nil
	assert.Nil(t, err)

	assert.True(t, len(result.Result) > 0)
}
