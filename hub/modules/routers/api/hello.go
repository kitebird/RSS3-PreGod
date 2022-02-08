package api

import (
	"net/http"

	"github.com/NaturalSelectionLabs/RSS3-PreGod/hub/modules/status"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/hub/modules/web"
	"github.com/gin-gonic/gin"
)

func GetHello(c *gin.Context) {
	w := web.Gin{C: c}

	w.JSONResponse(http.StatusOK, status.SUCCESS, gin.H{
		"hello": "world",
	})
}
