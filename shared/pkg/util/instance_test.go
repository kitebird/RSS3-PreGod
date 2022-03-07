package util_test

import (
	"testing"

	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/util"
	"github.com/stretchr/testify/assert"
)

func TestNewInstance(t *testing.T) {
	t.Parallel()

	instance, err := util.NewInstance("account", "0xC8b960D09C0078c18Dcbe7eB9AB9d816BcCa8944", "ethereum")
	assert.Nil(t, err)
	assert.Equal(t, instance.String(), "account:0xC8b960D09C0078c18Dcbe7eB9AB9d816BcCa8944@ethereum")

	_, err = util.NewInstance("foobar", "0xC8b960D09C0078c18Dcbe7eB9AB9d816BcCa8944", "ethereum")
	assert.Equal(t, err, util.ErrInvalidPrefix)
}

func BenchmarkNewInstance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = util.NewInstance("account", "0xC8b960D09C0078c18Dcbe7eB9AB9d816BcCa8944", "ethereum")
	}
}

func TestParseInstance(t *testing.T) {
	t.Parallel()

	instance, err := util.ParseInstance("account:0xC8b960D09C0078c18Dcbe7eB9AB9d816BcCa8944@ethereum")
	assert.Nil(t, err, err)
	assert.Equal(t, instance.String(), "account:0xC8b960D09C0078c18Dcbe7eB9AB9d816BcCa8944@ethereum")

	_, err = util.ParseInstance("foobar:0xC8b960D09C0078c18Dcbe7eB9AB9d816BcCa8944@ethereum")
	assert.ErrorIs(t, err, util.ErrInvalidPrefix)

	_, err = util.ParseInstance("account@ethereum")
	assert.ErrorIs(t, err, util.ErrInvalidIdentity)

	_, err = util.ParseInstance("account:NaturalSelectionLabs@gitlab")
	assert.ErrorIs(t, err, util.ErrInvalidPlatform)
}

func BenchmarkParseInstance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = util.ParseInstance("account:0xC8b960D09C0078c18Dcbe7eB9AB9d816BcCa8944@ethereum")
	}
}
