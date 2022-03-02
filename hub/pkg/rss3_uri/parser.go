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

// Parses the uri and returns the authority struct.
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
