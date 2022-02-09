package api

import (
	"net/http"

	"github.com/NaturalSelectionLabs/RSS3-PreGod/hub/pkg/status"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/hub/pkg/web"
	"github.com/gin-gonic/gin"
)

func GetHello(c *gin.Context) {
	w := web.Gin{C: c}

	w.JSONResponse(http.StatusOK, status.SUCCESS, gin.H{
		"hello": "world",
	})
}
