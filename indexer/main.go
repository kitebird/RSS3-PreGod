package main

import (
	"log"
	"time"

	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/api/arweave"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/api/gitcoin"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/crawler"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/processor"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/processor/item_stroge_task"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/router"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/cache"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/config"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/logger"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/rss3uri"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/web"
	jsoniter "github.com/json-iterator/go"
)

var jsoni = jsoniter.ConfigCompatibleWithStandardLibrary

func init() {
	if err := config.Setup(); err != nil {
		log.Fatalf("config.Setup err: %v", err)
	}

	if err := logger.Setup(); err != nil {
		log.Fatalf("config.Setup err: %v", err)
	}

	if err := cache.Setup(); err != nil {
		log.Fatalf("cache.Setup err: %v", err)
	}

	if err := processor.Setup(); err != nil {
		log.Fatalf("processor.Setup err: %v", err)
	}

}

func dispatchTasks(q chan processor.ProcessTaskHandler, ti time.Duration) {
	// TODO: Get all accounts
	instances := []rss3uri.PlatformInstance{}
	for _, i := range instances {
		for _, n := range i.Platform.ID().GetNetwork() {
			time.Sleep(ti)

			itemStrogeTask := item_stroge_task.NewItemStrogeParam(crawler.WorkParam{Identity: i.Identity, PlatformID: i.Platform.ID(), NetworkID: n})
			q <- itemStrogeTask.ProcessTaskHandler
		}
	}
}

func pollTasks(q chan processor.ProcessTaskHandler) {
	for {
		dispatchTasks(q, time.Minute)
		time.Sleep(24 * time.Hour)
	}
}

func main() {

	go pollTasks(processor.GlobalProcessor.LowQ)

	srv := &web.Server{
		RunMode:      config.Config.Indexer.Server.RunMode,
		HttpPort:     config.Config.Indexer.Server.HttpPort,
		ReadTimeout:  config.Config.Indexer.Server.ReadTimeout,
		WriteTimeout: config.Config.Indexer.Server.WriteTimeout,
		Handler:      router.InitRouter(),
	}

	srv.Start()

	// arweave crawler
	ar := arweave.NewArCrawler(
		1,
		500,
		10,
		2,
		600,
		"Ky1c1Kkt-jZ9sY1hvLF5nCf6WWdBhIU5Un_BMYh-t3c")
	ar.Start()

	// gitcoin crawler
	ethParam := gitcoin.NewParam(1, 10000, 10, 10, 600)
	polygonParam := gitcoin.NewParam(1, 10000, 10, 10, 600)
	zkParam := gitcoin.NewParam(1, 10000, 10, 10, 600)
	gc := gitcoin.NewGitcoinCrawler(ethParam, polygonParam, zkParam)

	go gc.PolygonStart()
	go gc.EthStart()
	go gc.ZkStart()

	defer logger.Logger.Sync()
}
