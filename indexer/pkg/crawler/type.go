package crawler

import (
	"time"

	jsoniter "github.com/json-iterator/go"

	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/db/model"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/constants"
)

type Crawler interface {
	Work(param WorkParam) error
	// GetResult return &{Assets, Notes, Items}
	GetResult() *CrawlerResult
	// GetBio
	// Since some apps have multiple bios,
	// they need to be converted into json and then collectively transmitted
	GetUserBio(Identity string) (string, error)
}

type CrawlerResult struct {
	Assets []*model.ItemId
	Notes  []*model.ItemId
	Items  []*model.Item
}

type WorkParam struct {
	Identity   string
	NetworkID  constants.NetworkID
	PlatformID constants.PlatformID // optional
	Limit      int                  // optional, aka Count, limit the number of items to be crawled

	TimeStamp   time.Time // optional
	BlockHeight int64     // optional
}

// CrawlerResult inherits the function by default

func (cr *CrawlerResult) Work(param WorkParam) error {
	return nil
}

func (cr *CrawlerResult) GetResult() *CrawlerResult {
	return cr
}

func (cr *CrawlerResult) GetUserBio(Identity string) (string, error) {
	return "", nil
}

type UserBios struct {
	Bios []string `json:"bios"`
}

func GetUserBioJson(bios []string) (string, error) {
	jsoni := jsoniter.ConfigCompatibleWithStandardLibrary

	userbios := UserBios{Bios: bios}
	userBioJson, err := jsoni.MarshalToString(userbios)

	if err != nil {
		return "", err
	}

	return userBioJson, nil
}
