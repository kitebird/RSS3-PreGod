package api

import (
	"net/http"

	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/crawler"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/processor"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/indexer/pkg/processor/user_bio_stroge_task"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/constants"
	"github.com/gin-gonic/gin"
)

type GetBioRequest struct {
	Identity   string               `form:"proof" binding:"required"`
	PlatformId constants.PlatformID `form:"platform_id" binding:"required"`
}

func GetBioHandlerFunc(c *gin.Context) {
	request := GetItemRequest{}
	if err := c.ShouldBind(&request); err != nil {
		return
	}

	// TODO Query data
	userBioStrogeTask := user_bio_stroge_task.NewUserBioStrogeTask(crawler.WorkParam{xx, xx, xx})
	processor.GlobalProcessor.UrgentQ <- userBioStrogeTask
	result := <-user_bio_stroge_task.ResultQ

	c.JSON(http.StatusOK, request)
}
