package api

import (
	"net/http"

	"github.com/NaturalSelectionLabs/RSS3-PreGod/hub/pkg/rss3_uri"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/hub/pkg/status"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/hub/pkg/web"
	"github.com/gin-gonic/gin"
)

type Uri struct {
	Authority string `uri:"authority" binding:"required"`
}

func GetInstance(c *gin.Context) {
	w := web.Gin{C: c}

	var uri Uri
	if err := c.ShouldBindUri(&uri); err != nil {
		w.JSONResponse(http.StatusBadRequest, status.INVALID_PARAMS, "invalid uri")

		return
	}

	authority, err := rss3_uri.ParseAuthority(uri.Authority)
	if err != nil {
		w.JSONResponse(http.StatusBadRequest, status.INVALID_PARAMS, "invalid uri: "+err.Error())

		return
	}

	// TODO: get instance from db

	w.JSONResponse(http.StatusOK, status.SUCCESS, gin.H{
		"authority": authority,
	})
}
