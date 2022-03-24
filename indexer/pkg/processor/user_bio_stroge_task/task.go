package user_bio_stroge_task

import (
	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/crawler"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/processor"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/logger"
)

type UserBioStrogeTask struct {
	processor.ProcessTaskParam
	ResultQ chan *UserBioStrogeResult
}

type UserBioStrogeResult struct {
	processor.ProcessTaskResult

	UserBio string
}

func NewUserBioStrogeTask(workParam crawler.WorkParam) *UserBioStrogeTask {
	return &UserBioStrogeTask{
		processor.ProcessTaskParam{
			TaskType:  processor.ProcessTaskTypeUserBioStroge,
			WorkParam: workParam,
		},
		make(chan *UserBioStrogeResult),
	}
}

func NewUserBioStrogeResult() *UserBioStrogeResult {
	return &UserBioStrogeResult{
		processor.ProcessTaskResult{
			TaskType:   processor.ProcessTaskTypeUserBioStroge,
			TaskResult: processor.ProcessTaskErrorCodeSuccess,
		},

		"",
	}
}

func (pt *UserBioStrogeTask) Fun() error {
	var err error

	var c crawler.Crawler

	var userBio string

	result := NewUserBioStrogeResult()

	c = processor.MakeCrawlers(pt.WorkParam.NetworkID)
	if c == nil {
		result.TaskResult = processor.ProcessTaskErrorCodeNotSupportedNetwork

		logger.Errorf("unsupported network id: %d", pt.WorkParam.NetworkID)

		goto RETURN
	}

	logger.Infof("c:%v", &c)

	userBio, err = c.GetUserBio(pt.WorkParam.Identity)

	if err != nil {
		result.TaskResult = processor.ProcessTaskErrorCodeNotFoundData

		goto RETURN
	}

	if len(userBio) > 0 {
		// TODO：add userbio into redis
		// redis.SetUserBio(userBio)
		// ctx := context.Background()
		// key := fmt.Sprintf("%s_%s_%s", pt.WorkParam.Identity,
		// 	pt.WorkParam.PlatformID.Symbol(),
		// )
		// cache.Set(ctx, key, userBio, 2)
		result.UserBio = userBio
	}

RETURN:
	pt.ResultQ <- result

	if err != nil {
		logger.Error(err)

		return err
	} else {
		return nil
	}
}
