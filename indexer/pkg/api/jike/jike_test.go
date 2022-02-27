package jike_test

import (
	"testing"

	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/api/jike"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	t.Parallel()

	// fmt.Printf("Jike config: %v\n", config.ThirdPartyConfig.Jike)

	err := jike.Login()

	assert.Nil(t, err)

	t.Log(jike.RefreshToken)
	t.Log(jike.AccessToken)

	assert.True(t,
		len(jike.RefreshToken) > 0)
}
