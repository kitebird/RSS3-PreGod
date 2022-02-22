package arweave

import (
	"fmt"

	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/util"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/types"
	"github.com/valyala/fastjson"
)

const arweaveEndpoint string = "https://arweave.net"
const arweaveGraphqlEndpoint string = "https://arweave.net/graphql"
const mirrorHost = "https://mirror.xyz/"

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

func ParseGraphqlNode(node string) (types.MirrorArticle, error) {
	var parser fastjson.Parser

	parsedJson, err := parser.Parse(node)
	if err != nil {
		return types.MirrorArticle{}, err
	}

	article := new(types.MirrorArticle)

	tags := parsedJson.GetArray("node", "tags")
	for _, tag := range tags {
		name := string(tag.GetStringBytes("name"))
		value := string(tag.GetStringBytes("value"))

		switch name {
		case "Contributor":
			article.Author = value
		case "Content-Digest":
			article.Digest = value
		case "Original-Content-Digest":
			article.OriginalDigest = value
		}

		article.Link = mirrorHost + "/" + article.Author + "/" + article.OriginalDigest
	}

	id := parsedJson.GetStringBytes("node", "id")

	content, err := GetContentByTxHash(string(id))
	if err != nil {
		return *article, err
	}

	parsedJson, err = parser.Parse(string(content))
	if err != nil {
		return *article, err
	}

	// title
	article.Title = string(parsedJson.GetStringBytes("content", "title"))
	// timestamp
	article.TimeStamp = parsedJson.GetUint64("content", "timestamp")
	// content
	article.Content = string(parsedJson.GetStringBytes("content", "body")) // timestamp

	return *article, nil
}
