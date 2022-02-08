package routers

import (
	"github.com/NaturalSelectionLabs/RSS3-PreGod/hub/modules/routers/api"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.GET("/hello", api.GetHello)

	return r
}
