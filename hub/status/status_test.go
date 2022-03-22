package status_test

import (
	"testing"

	"github.com/NaturalSelectionLabs/RSS3-PreGod/hub/status"
)

func TestName(t *testing.T) {
	t.Parallel()

	t.Log(status.CodeFileFieldError.Message())
}
