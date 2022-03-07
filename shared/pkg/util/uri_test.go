package util_test

import (
	"testing"

	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/constants"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/util"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Parallel()

	instance, err := util.NewInstance("account", "0xC8b960D09C0078c18Dcbe7eB9AB9d816BcCa8944", "ethereum")
	assert.Nil(t, err)

	uri := util.New(instance)
	assert.Equal(t, uri.String(), "rss3://account:0xC8b960D09C0078c18Dcbe7eB9AB9d816BcCa8944@ethereum")

	uri = util.New(&util.PlatformInstance{
		Prefix:   constants.PrefixNameAccount,
		Identity: "0xC8b960D09C0078c18Dcbe7eB9AB9d816BcCa8944",
		Platform: constants.PlatformSymbolEthereum,
	})
	assert.Equal(t, uri.String(), "rss3://account:0xC8b960D09C0078c18Dcbe7eB9AB9d816BcCa8944@ethereum")
}

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = util.NewInstance("account", "0xC8b960D09C0078c18Dcbe7eB9AB9d816BcCa8944", "ethereum")
	}
}

func TestParse(t *testing.T) {
	t.Parallel()

	uri, err := util.Parse("rss3://account:0xC8b960D09C0078c18Dcbe7eB9AB9d816BcCa8944@ethereum")
	assert.Nil(t, err, err)
	assert.Equal(t, uri.String(), "rss3://account:0xC8b960D09C0078c18Dcbe7eB9AB9d816BcCa8944@ethereum")

	_, err = util.Parse("https://github.com/NaturalSelectionLabs/RSS3-PreGod")
	assert.ErrorIs(t, err, util.ErrInvalidScheme)

	_, err = util.Parse("rss3://foobar:0xC8b960D09C0078c18Dcbe7eB9AB9d816BcCa8944@ethereum")
	assert.ErrorIs(t, err, util.ErrInvalidPrefix)

	_, err = util.Parse("rss3://account:@ethereum")
	assert.ErrorIs(t, err, util.ErrInvalidIdentity)

	_, err = util.Parse("rss3://account:NaturalSelectionLabs@gitlab")
	assert.ErrorIs(t, err, util.ErrInvalidPlatform)
}

func BenchmarkParse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = util.Parse("rss3://account:0xC8b960D09C0078c18Dcbe7eB9AB9d816BcCa8944@ethereum")
	}
}
