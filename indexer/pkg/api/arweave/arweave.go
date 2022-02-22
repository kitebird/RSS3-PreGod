package arweave

import (
	"fmt"

	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/util"
)

const endpoint = "https://arweave.net"

// GetContentByTxHash gets transaction content by tx hash.
func GetContentByTxHash(hash string) ([]byte, error) {
	var headers = map[string]string{
		"Origin":  "https://viewblock.io",
		"Referer": "https://viewblock.io",
	}

	url := fmt.Sprintf("%s/%s", endpoint, hash)

	return util.Get(url, headers)
}

// GetTransacitons gets all transactions using filters.
func GetTransacitons(from, to uint64, owner string) ([]byte, error) {
	var headers = map[string]string{
		"Accept-Encoding": "gzip, deflate, br",
		"Content-Type":    "application/json",
		"Accept":          "application/json",
		"Connection":      "keep-alive",
		"DNT":             "1",
		"Origin":          "https://arweave.net",
	}

	queryVariables :=
		"{\"query\":\"query { transactions( " +
			"block: { min: %d, max: %d } " +
			"owners: [\\\"%s\\\"] " +
			"sort: HEIGHT_ASC ) { edges { node {id tags { name value } } } }" +
			"}\"}"
	data := fmt.Sprintf(queryVariables, from, to, owner)
	url := fmt.Sprintf("%s/graphql", endpoint)

	return util.Post(url, headers, data)
}
