package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetNoteListRequest struct {
	Instance string `uri:"instance" binding:"required"`
	NoteID   string `uri:"note_id" binding:"required"`
}

func GetNoteListRequestHandlerFunc(c *gin.Context) {
	request := GetNoteListRequest{}
	if err := c.ShouldBindUri(&request); err != nil {
		return
	}

	// TODO Query data from database
	// TODO No test data available
	c.JSON(http.StatusOK, request)
}
