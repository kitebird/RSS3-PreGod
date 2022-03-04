package rss3_uri

import (
	"errors"
	"strings"

	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/constants"
)

type Authority struct {
	Prefix   constants.PrefixName   `json:"prefix"`
	Identity string                 `json:"identity"`
	Platform constants.PlatformName `json:"platform"`
}

// ParseAuthority parses the uri and returns the authority struct.
// Returns error if the uri is invalid.
func ParseAuthority(uri string) (*Authority, error) {
	s := strings.SplitN(uri, "@", 2)

	if len(s) != 2 {
		return nil, errors.New("no platform name")
	}

	pi := strings.SplitN(s[0], ":", 2)

	if len(pi) != 2 {
		return nil, errors.New("no prefix")
	}

	prefix := pi[0]
	identity := pi[1]
	platform := s[1]

	if !constants.IsValidPrefix(prefix) {
		return nil, errors.New("invalid prefix")
	}

	if !constants.IsValidPlatformName(platform) {
		return nil, errors.New("invalid platform name")
	}

	return &Authority{
		Prefix:   constants.PrefixName(prefix),
		Identity: identity,
		Platform: constants.PlatformName(platform),
	}, nil
}

// AccountInstanceURI returns an account instance URI
// example:
// rss3://account:0xC8b960D09C0078c18Dcbe7eB9AB9d816BcCa8944@evm
// rss3://account:DIYgod@twitter
func AccountInstanceURI(identity string, accountPlatform constants.PlatformName) string {
	return string(constants.Prefix_RSS3) + string(constants.Prefix_Account+":") +
		identity + "@" + string(accountPlatform)
}

// ItemInstanceURI returns an automatically generated Instance URI for Item.
// example:
// rss3://asset:ethereum-0xacbe98efe2d4d103e221e04c76d7c55db15c8e89-5@evm
func ItemInstanceURI(itemType constants.Prefix, uniqueId string, itemPlatform constants.PlatformName) string {
	return string(constants.Prefix_RSS3+itemType) + ":" + uniqueId + "@" + string(itemPlatform)
}

// AssetInstanceURI returns an asset instance URI
// example:
// rss3://asset:0xb9619cf4f875cdf0e3ce48b28a1c725bc4f6c0fb-1024@ethereum
func AssetInstanceURI(assetAddress, tokenId string, assetPlatform constants.PlatformName) string {
	return string(constants.Prefix_RSS3) + string(constants.Prefix_Asset+":") + assetAddress + "-" + tokenId + "@" + string(assetPlatform)
}

// NoteInstanceURI returns an note instanceURI
// example:
// rss3://note:5591079b-1f5b-4ae9-8209-51b18f0d3be0@twitter
func NoteInstanceURI(noteUUID string, itemPlatform constants.PlatformName) string {
	return string(constants.Prefix_RSS3) + string(constants.Prefix_Note+":") + noteUUID + "@" + string(itemPlatform)
}

// ItemURI returns an itemURI
// example:
// rss3://account:0xC8b960D09C0078c18Dcbe7eB9AB9d816BcCa8944@evm/notes/5591079b-1f5b-4ae9-8209-51b18f0d3be0
func ItemURI(instanceId, itemType, itemUUID string) string {
	return string(constants.Prefix_RSS3) + instanceId + "/" + itemType + "/" + itemUUID
}
