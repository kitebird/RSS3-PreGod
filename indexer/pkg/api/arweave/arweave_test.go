package arweave_test

import (
	"testing"

	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/api/arweave"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fastjson"
)

func TestGetContentByTxHash(t *testing.T) {
	t.Parallel()

	hash := "BhM-D1bsQkaqi72EEG1aRVs4Nv5bZZIW-mH8yEdIDWA"
	response, err := arweave.GetContentByTxHash(hash)
	// assert for nil
	assert.Nil(t, err)
	assert.True(t, len(response) > 0)

	var parser fastjson.Parser
	parsedJson, parseErr := parser.Parse(string(response))
	assert.Nil(t, parseErr)

	// check title
	title := parsedJson.GetStringBytes("content", "title")
	assert.True(t, len(title) > 0)
	// check body
	body := parsedJson.GetStringBytes("content", "body")
	assert.True(t, len(body) > 0)
	// check contributor
	contributor := parsedJson.GetStringBytes("authorship", "contributor")
	assert.True(t, len(contributor) > 0)
	// check originalDigest
	originalDigest := parsedJson.GetStringBytes("originalDigest")
	assert.True(t, len(originalDigest) > 0)
}

func TestGetTransacitons(t *testing.T) {
	t.Parallel()

	owner := "Ky1c1Kkt-jZ9sY1hvLF5nCf6WWdBhIU5Un_BMYh-t3c"
	response, err := arweave.GetTransactions(877250, 877250, owner)
	// assert for nil
	assert.Nil(t, err)
	assert.True(t, len(response) > 0)

	var parser fastjson.Parser
	parsedJson, parseErr := parser.Parse(string(response))
	assert.Nil(t, parseErr)

	edges := parsedJson.GetArray("data", "transactions", "edges")
	assert.True(t, len(edges) > 0)

	// edges
	for _, edge := range edges {
		id := edge.GetStringBytes("node", "id")
		assert.True(t, len(id) > 0)
	}
}
