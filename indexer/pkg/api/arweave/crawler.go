package arweave

import (
	"errors"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/db"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/db/model"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/constants"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/logger"
)

var ErrTimeout = errors.New("received timeout")
var ErrInterrupt = errors.New("received interrupt")

type crawlConfig struct {
	fromHeight    int64
	confirmations int64
	step          int64
	sleepInterval time.Duration
}
type arCrawler struct {
	identity  ArAccount
	interrupt chan os.Signal
	complete  chan error
	cfg       *crawlConfig
}

func NewArCrawler(identity ArAccount, crawlCfg *crawlConfig) *arCrawler {
	return &arCrawler{
		identity,
		make(chan os.Signal, 1),
		make(chan error),
		crawlCfg,
	}
}

func (ar *arCrawler) run() error {
	startBlockHeight := ar.cfg.fromHeight
	step := ar.cfg.step
	tempDelay := ar.cfg.sleepInterval

	latestConfirmedBlockHeight, err := GetLatestBlockHeightWithConfirmations(ar.cfg.confirmations)
	if err != nil {
		return err
	}

	for {
		// handle interrupt
		if ar.gotInterrupt() {
			return ErrInterrupt
		}

		endBlockHeight := startBlockHeight + step
		if latestConfirmedBlockHeight <= endBlockHeight {
			time.Sleep(tempDelay)

			latestConfirmedBlockHeight, err = GetLatestBlockHeightWithConfirmations(ar.cfg.confirmations)
			if err != nil {
				return err
			}

			step = DefaultCrawlStep
		} else {
			step = ar.cfg.step
		}

		//TODO: Sleep here
		ar.getArticles(startBlockHeight, latestConfirmedBlockHeight, ar.identity)
	}
}

// TODO: make it parseMirror args
func (ar *arCrawler) getArticles(from, to int64, owner ArAccount) error {
	articles, err := GetArticles(from, to, owner)
	if err != nil {
		return err
	}

	items := make([]*model.Item, 0)

	for _, article := range articles {
		attachment := model.Attachment{
			Type:     "body",
			Content:  article.Content,
			MimeType: "text/markdown",
		}

		tsp, err := time.Parse(time.RFC3339, strconv.FormatInt(article.TimeStamp, 10))
		if err != nil {
			//TODO: may send to a error queue or whatever in the future
			logger.Error(err)

			tsp = time.Now()
		}

		ni := model.NewItem(
			constants.NetworkSymbolArweaveMainnet.GetID(),
			article.TxHash,
			nil,
			[]string{"https://arweave.net/" + article.TxHash, "https://mirror.xyz/" + article.Author + "/" + article.OriginalDigest},
			constants.ItemTagsMirrorEntry,
			[]string{article.Author},
			article.Title,
			article.Content, // TODO: According to RIP4, if the body is too long, then only record part of the body, followed by ... at the end
			[]model.Attachment{attachment},
			tsp,
		)

		items = append(items, ni)
	}

	setDB(items)

	return nil
}

func setDB(items []*model.Item) {
	for _, item := range items {
		db.InsertItem(item)
	}
}

func (ar *arCrawler) Start() error {
	signal.Notify(ar.interrupt, os.Interrupt)

	go func() {
		ar.complete <- ar.run()
	}()

	select {
	case err := <-ar.complete:
		return err
	default:
		return nil
	}
}

func (ar *arCrawler) gotInterrupt() bool {
	select {
	case <-ar.interrupt:
		signal.Stop(ar.interrupt)

		return true
	default:
		return false
	}
}
