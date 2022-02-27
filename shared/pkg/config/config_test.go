package config_test

import (
	"testing"

	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestSetup(t *testing.T) {
	t.Parallel()

	err := config.Setup()
	assert.Nil(t, err)
}
