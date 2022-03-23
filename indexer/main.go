package main

import (
	"log"
	"time"

	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/crawler"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/db"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/processor"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/router"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/cache"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/config"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/logger"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/rss3uri"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/web"
)

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

func dispatchTasks(q chan *processor.ProcessTaskParam, ti time.Duration) {
	// TODO: Get all accounts
	instances := []rss3uri.PlatformInstance{}
	for _, i := range instances {
		for _, n := range i.Platform.ID().GetNetwork() {
			time.Sleep(ti)

			itemStrogeTask := item_stroge_task.NewItemStrogeParam(crawler.WorkParam{Identity: i.Identity, PlatformID: i.Platform.ID(), NetworkID: n})
			q <- &itemStrogeTask.ProcessTaskParam
		}
	}
}

func pollTasks(q chan *processor.ProcessTaskParam) {
	for {
		dispatchTasks(q, time.Minute)
		time.Sleep(24 * time.Hour)
	}
}

func main() {
	urgentQ := processor.NewTaskQueue()
	lowQ := processor.NewTaskQueue()
	highQ := processor.NewTaskQueue()

	// 这里待会要考虑一下封装
	w := processor.NewProcessor(urgentQ, lowQ, highQ)
	go w.ListenAndServe()

	// TODO: listen tasks from mq
	// TODO: gracefully exit
	go pollTasks(lowQ)

	srv := &web.Server{
		RunMode:      config.Config.Indexer.Server.RunMode,
		HttpPort:     config.Config.Indexer.Server.HttpPort,
		ReadTimeout:  config.Config.Indexer.Server.ReadTimeout,
		WriteTimeout: config.Config.Indexer.Server.WriteTimeout,
		Handler:      router.InitRouter(),
	}

	addr := srv.Start()

	logger.Infof("Start http server listening on http://%s", addr)
	defer logger.Logger.Sync()
}
