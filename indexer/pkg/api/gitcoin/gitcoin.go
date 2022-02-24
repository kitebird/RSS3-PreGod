package gitcoin

import (
	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/util"
)

const grantUrl = "https://gitcoin.co/grants/grants.json"

// GetGrants returns all grant projects.
func GetGrants() (content []byte, err error) {
	content, err = util.Get(grantUrl, nil)

	return
}
