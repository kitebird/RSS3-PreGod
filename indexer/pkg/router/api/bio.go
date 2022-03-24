package api

import (
	"net/http"

	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/crawler"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/processor"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/processor/user_bio_stroge_task"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/constants"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/logger"
	"github.com/gin-gonic/gin"
)

type GetBioRequest struct {
	Identity   string               `form:"proof" binding:"required"`
	PlatformId constants.PlatformID `form:"platform_id" binding:"required"`
}

// struct GetBioResponse struct {

// }

var (
	// Since the transmitted parameter is only PlatformId
	// Currently, the platform and network for pulling bio are the same
	// , so there is a need for a place to transfer to each other.
	platform2Network = map[constants.PlatformID]constants.NetworkID{
		constants.PlatformIDTwitter: constants.NetworkIDTwitter,
		constants.PlatformIDJike:    constants.NetworkIDJike,
		constants.PlatformIDMisskey: constants.NetworkIDMisskey,
	}
)

func GetBioHandlerFunc(c *gin.Context) {
	request := GetBioRequest{}
	if err := c.ShouldBind(&request); err != nil {
		logger.Errorf("%s", err.Error())
		return
	}

	if len(request.Identity) > 0 || !constants.IsValidPlatformSymbol(request.PlatformId) {

	}

	userBioStrogeTask := user_bio_stroge_task.NewUserBioStrogeTask(
		crawler.WorkParam{
			Identity:   request.Identity,
			PlatformID: request.PlatformId,
			NetworkID:  platform2Network[request.PlatformId],
		})
	processor.GlobalProcessor.UrgentQ <- userBioStrogeTask
	result := <-userBioStrogeTask.ResultQ

	c.JSON(http.StatusOK, request)
}
