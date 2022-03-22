package user_bio_stroge_task

import (
	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/crawler"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/processor"
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
	var c crawler.Crawler

	userBio, err := c.GetUserBio(pt.WorkParam)
	result := NewUserBioStrogeResult()

	if err != nil {
		result.TaskResult = processor.ProcessTaskErrorCodeFoundData
		ResultQ <- result

		return err
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

	ResultQ <- result

	return nil
}
