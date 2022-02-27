package poap_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetActions(t *testing.T) {
	t.Parallel()

	result, err := GetActions("")

	assert.Nil(t, err)

	assert.True(t, len(result) > 0)
}
