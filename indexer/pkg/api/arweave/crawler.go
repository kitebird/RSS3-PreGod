package arweave

import (
	"strconv"
	"time"

	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/crawler"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/db/model"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/constants"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/logger"
)

type arCrawler struct {
	crawler.CrawlerResult
	fromHeight int64
	toHeight   int64
}

func NewArCrawler() crawler.Crawler {
	return &arCrawler{
		crawler.CrawlerResult{
			Assets: []*model.ItemId{},
			Notes:  []*model.ItemId{},
			Items:  []*model.Item{},
		},
		0,
		0,
	}
}

func (ar *arCrawler) Work(param crawler.WorkParam) error {
	startBlockHeight := int64(1)
	step := param.Step
	tempDelay := param.SleepInterval

	for {
		latestBlockHeight, err := GetLatestBlockHeight()
		if err != nil {
			return err
		}

		endBlockHeight := startBlockHeight + step
		if latestBlockHeight <= endBlockHeight {
			time.Sleep(tempDelay)

			latestBlockHeight, err = GetLatestBlockHeight()
			if err != nil {
				return err
			}

			endBlockHeight = latestBlockHeight
			step = 10
		}

		ar.getArticles(startBlockHeight, endBlockHeight, param.Identity)
	}
}

func (ar *arCrawler) getArticles(from, to int64, owner string) error {
	articles, err := GetArticles(from, to, owner)
	if err != nil {
		return err
	}

	for _, article := range articles {
		attachment := model.Attachment{
			Type:     "body",
			Content:  article.Content,
			MimeType: "text/markdown",
		}

		tsp, err := time.Parse(time.RFC3339, strconv.FormatInt(article.TimeStamp, 10))
		if err != nil {
			logger.Error(err)

			tsp = time.Now()
		}

		ni := model.NewItem(
			constants.NetworkSymbolArweaveMainnet.GetID(),
			article.Digest,
			model.Metadata{
				"network": constants.NetworkSymbolArweaveMainnet,
				"proof":   article.Digest,
			},
			constants.ItemTagsMirrorEntry,
			[]string{article.Author},
			article.Title,
			article.Content, // TODO: According to RIP4, if the body is too long, then only record part of the body, followed by ... at the end
			[]model.Attachment{attachment},
			tsp,
		)

		ar.Items = append(ar.Items, ni)
	}

	return nil
}

func (ar *arCrawler) GetResult() *crawler.CrawlerResult {
	return &crawler.CrawlerResult{
		Assets: ar.Assets,
		Notes:  ar.Notes,
		Items:  ar.Items,
	}
}
