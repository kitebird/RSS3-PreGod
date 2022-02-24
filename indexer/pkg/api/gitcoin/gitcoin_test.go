package gitcoin_test

import (
	"testing"

	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/api/gitcoin"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fastjson"
)

func TestGetGrants(t *testing.T) {
	t.Parallel()

	res, err := gitcoin.GetGrants()
	assert.Nil(t, err)
	assert.NotEmpty(t, res)

	var parser fastjson.Parser
	parsedJson, parseErr := parser.Parse(string(res))
	assert.Nil(t, parseErr)

	arr := parsedJson.GetArray()
	for _, v := range arr {
		item := v.GetArray()
		if len(item) == 2 && item[1].String() != "\"0x0\"" {
			// check title
			title := item[0].String()
			assert.NotEmpty(t, title)
			// check address
			adminAddress := item[1].String()
			assert.NotEmpty(t, adminAddress)
		}
	}
}
