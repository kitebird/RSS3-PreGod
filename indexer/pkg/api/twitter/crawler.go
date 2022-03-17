package twitter

import (
	"fmt"
	"time"

	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/crawler"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/db/model"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/constants"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/logger"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/rss3uri"
)

type twitterCrawler struct {
	crawler.CrawlerResult
}

func NewTwitterCrawler() crawler.Crawler {
	return &twitterCrawler{
		crawler.CrawlerResult{
			Items: []*model.Item{},
			Notes: []*model.ItemId{},
		},
	}
}

const DefaultTwitterCount = 200

func (tc *twitterCrawler) Work(param crawler.WorkParam) error {
	if param.NetworkID != constants.NetworkIDTwitter {
		return fmt.Errorf("network is not twitter")
	}

	networkSymbol := constants.NetworkSymbolTwitter

	networkId := networkSymbol.GetID()

	contentInfos, err := GetTimeline(param.Identity, DefaultTwitterCount)
	if err != nil {
		return err
	}

	author, err := rss3uri.NewInstance("account", param.Identity, string(constants.PlatformSymbolTwitter))
	if err != nil {
		return err
	}

	for _, contentInfo := range contentInfos {
		tsp, err := contentInfo.GetTsp()
		if err != nil {
			// TODO: log error
			logger.Error(tsp, err)
			tsp = time.Now()
		}

		ni := model.NewItem(
			networkId,
			"",
			model.Metadata{},
			constants.ItemTagsTweet,
			[]string{author.String()},
			"",
			contentInfo.PreContent,
			[]model.Attachment{},
			tsp,
		)

		tc.Items = append(tc.Items, ni)
		tc.Notes = append(tc.Notes, &model.ItemId{
			NetworkID: networkId,
			Proof:     "",
		})
	}

	return nil
}

func (tc *twitterCrawler) GetResult() *crawler.CrawlerResult {
	return &crawler.CrawlerResult{
		Assets: tc.Assets,
		Notes:  tc.Notes,
	}
}

func (tc *twitterCrawler) GetUserBio(param crawler.WorkParam) (string, error) {
	userShow, err := GetUserShow(param.Identity)

	if err != nil {
		return "", err
	}

	return userShow.Description, nil
}
