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

func TestGetProject(t *testing.T) {
	t.Parallel()

	adminAddress := "0xf634ec94939efd57cb888fa8451c1e0d0f973c23"
	res, err := gitcoin.GetProject(adminAddress)
	assert.Nil(t, err)
	assert.NotEmpty(t, res)
}

func TestGetGrantsInfo(t *testing.T) {
	t.Parallel()

	res, err := gitcoin.GetGrantsInfo()
	assert.Nil(t, err)
	assert.NotEmpty(t, res)
}

func TestGetProjectsInfo(t *testing.T) {
	t.Parallel()

	res, err := gitcoin.GetProjectsInfo("0x8c23B96f2fb77AaE1ac2832debEE30f09da7af3C", "RSS3")
	assert.Nil(t, err)
	assert.NotEmpty(t, res)
}

func TestGetDonations(t *testing.T) {
	t.Parallel()

	_, err := gitcoin.GetDonations(12605342, 12605343)
	assert.Nil(t, err)
	//assert.NotEmpty(t, res)
}
