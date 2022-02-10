package api

import (
	"net/http"
	"strings"

	"github.com/NaturalSelectionLabs/RSS3-PreGod/hub/pkg/status"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/hub/pkg/web"
	"github.com/gin-gonic/gin"
)

type Uri struct {
	Instance string `uri:"instance" binding:"required"`
}

func GetInstance(c *gin.Context) {
	w := web.Gin{C: c}

	var uri Uri
	if err := c.ShouldBindUri(&uri); err != nil {
		w.JSONResponse(http.StatusBadRequest, status.INVALID_PARAMS, "invalid uri")
		return
	}

	s := strings.Split(uri.Instance, "@")
	if len(s) != 2 {
		w.JSONResponse(http.StatusBadRequest, status.INVALID_PARAMS, "invalid signable account")
		return
	}

	address := s[0]
	platform := s[1]

	// TODO: get instance from db

	w.JSONResponse(http.StatusOK, status.SUCCESS, gin.H{
		"address":  address,
		"platform": platform,
	})
}
