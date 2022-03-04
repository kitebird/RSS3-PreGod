package rss3uri_test

import (
	"github.com/NaturalSelectionLabs/RSS3-PreGod/hub/pkg/rss3_uri"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/hub/pkg/rss3uri"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewInstance(t *testing.T) {
	instance, err := rss3uri.NewInstance("account", "0xC8b960D09C0078c18Dcbe7eB9AB9d816BcCa8944", "evm")
	assert.Nil(t, err)
	assert.Equal(t, instance.String(), "account:0xC8b960D09C0078c18Dcbe7eB9AB9d816BcCa8944@evm")

	instance, err = rss3uri.NewInstance("foobar", "0xC8b960D09C0078c18Dcbe7eB9AB9d816BcCa8944", "evm")
	assert.Equal(t, err, rss3uri.ErrInvalidPrefix)
}

func BenchmarkNewInstance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = rss3uri.NewInstance("account", "0xC8b960D09C0078c18Dcbe7eB9AB9d816BcCa8944", "evm")
	}
}

func TestParseInstance(t *testing.T) {
	instance, err := rss3uri.ParseInstance("account:0xC8b960D09C0078c18Dcbe7eB9AB9d816BcCa8944@evm")
	assert.Nil(t, err, err)
	assert.Equal(t, instance.String(), "account:0xC8b960D09C0078c18Dcbe7eB9AB9d816BcCa8944@evm")

	instance, err = rss3uri.ParseInstance("foobar:0xC8b960D09C0078c18Dcbe7eB9AB9d816BcCa8944@evm")
	assert.ErrorIs(t, err, rss3uri.ErrInvalidPrefix)

	instance, err = rss3uri.ParseInstance("account@evm")
	assert.ErrorIs(t, err, rss3uri.ErrInvalidIdentity)

	instance, err = rss3uri.ParseInstance("account:NaturalSelectionLabs@github")
	assert.ErrorIs(t, err, rss3uri.ErrInvalidPlatform)
}

func BenchmarkParseInstanceOld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = rss3_uri.ParseAuthority("account:0xC8b960D09C0078c18Dcbe7eB9AB9d816BcCa8944@evm")
	}
}

func BenchmarkParseInstance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = rss3uri.ParseInstance("account:0xC8b960D09C0078c18Dcbe7eB9AB9d816BcCa8944@evm")
	}
}
