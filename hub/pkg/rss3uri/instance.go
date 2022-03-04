package rss3uri

import (
	"fmt"

	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/constants"
)

var (
	_ fmt.Stringer = &Instance{}
)

type Instance struct {
	Prefix   constants.PrefixName   `json:"prefix"`
	Identity string                 `json:"identity"`
	Platform constants.PlatformName `json:"platform"`
}

func (i *Instance) String() string {
	return fmt.Sprintf("%s:%s@%s", i.Prefix, i.Identity, i.Platform)
}

func NewInstance(prefix, identity, platform string) (*Instance, error) {
	if !constants.IsValidPrefix(prefix) {
		return nil, ErrInvalidPrefix
	}

	if identity == "" {
		return nil, ErrInvalidIdentity
	}

	if !constants.IsValidPlatformName(platform) {
		return nil, ErrInvalidPlatform
	}

	return &Instance{
		Prefix:   constants.PrefixName(prefix),
		Identity: identity,
		Platform: constants.PlatformName(platform),
	}, nil
}

func ParseInstance(rawInstance string) (*Instance, error) {
	uri, err := Parse(fmt.Sprintf("%s://%s", Scheme, rawInstance))
	if err != nil {
		return nil, err
	}

	return uri.Instance, nil
}
