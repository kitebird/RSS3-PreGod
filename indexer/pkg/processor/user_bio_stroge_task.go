package processor

import "github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/crawler"

type userBioStrogeTask struct {
	ProcessTaskParam
}

func NewUserBioStrogeTask() ProcessTaskUnit {
	return &userBioStrogeTask{
		ProcessTaskParam{
			TaskType: ProcessTaskTypeItemStroge,
		},
	}
}

func (pt *userBioStrogeTask) Fun() error {

	var c crawler.Crawler

	userBio, err := c.GetUserBio(pt.WorkParam)

	if err != nil {
		return err
	}

	if len(userBio) > 0 {
		// redis.SetUserBio(userBio)
	} else {
		// redis.SetUserBio(error)
	}

	return nil
}
