package arweave

import (
	"fmt"

	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/util"
)

const arweaveEndpoint string = "https://arweave.net"
const arweaveGraphqlEndpoint string = "https://arweave.net/graphql"

// GetContentByTxHash gets transaction content by tx hash.
func GetContentByTxHash(hash string) ([]byte, error) {
	var headers = map[string]string{
		"Origin":  "https://viewblock.io",
		"Referer": "https://viewblock.io",
	}

	return util.Get(arweaveEndpoint+"/"+hash, headers)
}

// GetTransactions gets all transactions using filters.
func GetTransactions(from, to uint64, owner string) ([]byte, error) {
	var headers = map[string]string{
		"Accept-Encoding": "gzip, deflate, br",
		"Content-Type":    "application/json",
		"Accept":          "application/json",
		"Origin":          "https://arweave.net",
	}

	queryVariables :=
		"{\"query\":\"query { transactions( " +
			"block: { min: %d, max: %d } " +
			"owners: [\\\"%s\\\"] " +
			"sort: HEIGHT_ASC ) { edges { node {id tags { name value } } } }" +
			"}\"}"
	data := fmt.Sprintf(queryVariables, from, to, owner)

	return util.Post(arweaveGraphqlEndpoint, headers, data)
}
