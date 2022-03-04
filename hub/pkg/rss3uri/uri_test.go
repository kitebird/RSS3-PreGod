package rss3uri_test

import (
	"testing"

	"github.com/NaturalSelectionLabs/RSS3-PreGod/hub/pkg/rss3uri"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/constants"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Parallel()

	instance, err := rss3uri.NewInstance("account", "0xC8b960D09C0078c18Dcbe7eB9AB9d816BcCa8944", "evm")
	assert.Nil(t, err)

	uri := rss3uri.New(instance)
	assert.Equal(t, uri.String(), "rss3://account:0xC8b960D09C0078c18Dcbe7eB9AB9d816BcCa8944@evm")

	uri = rss3uri.New(&rss3uri.Instance{
		Prefix:   constants.PrefixName_Account,
		Identity: "0xC8b960D09C0078c18Dcbe7eB9AB9d816BcCa8944",
		Platform: constants.PlatformName_Evm,
	})
	assert.Equal(t, uri.String(), "rss3://account:0xC8b960D09C0078c18Dcbe7eB9AB9d816BcCa8944@evm")
}

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = rss3uri.NewInstance("account", "0xC8b960D09C0078c18Dcbe7eB9AB9d816BcCa8944", "evm")
	}
}

func TestParse(t *testing.T) {
	t.Parallel()

	uri, err := rss3uri.Parse("rss3://account:0xC8b960D09C0078c18Dcbe7eB9AB9d816BcCa8944@evm")
	assert.Nil(t, err, err)
	assert.Equal(t, uri.String(), "rss3://account:0xC8b960D09C0078c18Dcbe7eB9AB9d816BcCa8944@evm")

	_, err = rss3uri.Parse("https://github.com/NaturalSelectionLabs/RSS3-PreGod")
	assert.ErrorIs(t, err, rss3uri.ErrInvalidScheme)

	_, err = rss3uri.Parse("rss3://foobar:0xC8b960D09C0078c18Dcbe7eB9AB9d816BcCa8944@evm")
	assert.ErrorIs(t, err, rss3uri.ErrInvalidPrefix)

	_, err = rss3uri.Parse("rss3://account:@evm")
	assert.ErrorIs(t, err, rss3uri.ErrInvalidIdentity)

	_, err = rss3uri.Parse("rss3://account:NaturalSelectionLabs@github")
	assert.ErrorIs(t, err, rss3uri.ErrInvalidPlatform)
}

func BenchmarkParse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = rss3uri.Parse("rss3://account:0xC8b960D09C0078c18Dcbe7eB9AB9d816BcCa8944@evm")
	}
}
