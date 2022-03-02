package rss3_uri_test

import (
	"testing"

	"github.com/NaturalSelectionLabs/RSS3-PreGod/hub/pkg/rss3_uri"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/constants"
	"github.com/stretchr/testify/assert"
)

func TestAccountInstanceURI(t *testing.T) {
	t.Parallel()

	uri := rss3_uri.AccountInstanceURI("0xC8b960D09C0078c18Dcbe7eB9AB9d816BcCa8944",
		constants.PlatformName_Evm)
	assert.Equal(t, "rss3://account:0xC8b960D09C0078c18Dcbe7eB9AB9d816BcCa8944@evm", uri)

	uri = rss3_uri.AccountInstanceURI("DIYgod", constants.PlatformName_Twitter)
	assert.Equal(t, "rss3://account:DIYgod@twitter", uri)
}

func TestItemInstanceURI(t *testing.T) {
	t.Parallel()

	uri := rss3_uri.ItemInstanceURI(constants.Prefix_Asset,
		"ethereum-0xacbe98efe2d4d103e221e04c76d7c55db15c8e89-5", constants.PlatformName_Evm)
	assert.Equal(t, "rss3://asset:ethereum-0xacbe98efe2d4d103e221e04c76d7c55db15c8e89-5@evm", uri)
}

func TestAssetInstanceURI(t *testing.T) {
	t.Parallel()

	uri := rss3_uri.AssetInstanceURI("0xb9619cf4f875cdf0e3ce48b28a1c725bc4f6c0fb",
		"1024", constants.PlatformName_Evm)
	assert.Equal(t, "rss3://asset:0xb9619cf4f875cdf0e3ce48b28a1c725bc4f6c0fb-1024@ethereum", uri)
}

func TestNoteInstanceURI(t *testing.T) {
	t.Parallel()

	uri := rss3_uri.NoteInstanceURI("5591079b-1f5b-4ae9-8209-51b18f0d3be0", constants.PlatformName_Twitter)
	assert.Equal(t, "rss3://note:5591079b-1f5b-4ae9-8209-51b18f0d3be0@twitter", uri)
}

func TestItemURI(t *testing.T) {
	t.Parallel()

	expected :=
		"rss3://account:0xC8b960D09C0078c18Dcbe7eB9AB9d816BcCa8944@ethereum/notes/5591079b-1f5b-4ae9-8209-51b18f0d3be0"
	uri := rss3_uri.ItemURI("account:0xC8b960D09C0078c18Dcbe7eB9AB9d816BcCa8944@ethereum",
		"notes", "5591079b-1f5b-4ae9-8209-51b18f0d3be0")
	assert.Equal(t, expected, uri)
}
