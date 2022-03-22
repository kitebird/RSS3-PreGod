package user_bio_stroge_task

import (
	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/crawler"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/processor"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/logger"
)

type UserBioStrogeTask struct {
	processor.ProcessTaskParam
}

type UserBioStrogeResult struct {
	processor.ProcessTaskResult

	UserBio string
}

var ResultQ = make(chan *UserBioStrogeResult)

func NewUserBioStrogeTask(workParam crawler.WorkParam) *UserBioStrogeTask {
	return &UserBioStrogeTask{
		processor.ProcessTaskParam{
			TaskType:  processor.ProcessTaskTypeUserBioStroge,
			WorkParam: workParam,
		},
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

	userBio, err = c.GetUserBio(pt.WorkParam)

	if err != nil {
		result.TaskResult = processor.ProcessTaskErrorCodeNotFoundData

		goto RETURN
	}

	if len(userBio) > 0 {
		// redis.SetUserBio(userBio)
		// ctx := context.Background()

		// key := fmt.Sprintf("%s_%s_%s", pt.WorkParam.Identity,
		// 	pt.WorkParam.PlatformID.Symbol(),
		// )

		// cache.Set(ctx, key, userBio, 2)
		result.UserBio = userBio
	}

RETURN:
	ResultQ <- result

	if err != nil {
		logger.Error(err)

		return err
	} else {
		return nil
	}
}
