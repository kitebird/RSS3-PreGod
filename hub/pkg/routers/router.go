package routers

import (
	"github.com/NaturalSelectionLabs/RSS3-PreGod/hub/pkg/routers/api"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.GET("/ping", api.Ping)

	// instance
	r.GET("/:signableAccount", api.GetInstance)

	// // items
	// r.GET("/:signableAccount/:itemType/:itemUuid", api.GetItem)
	// r.GET("/:signableAccount/list/:itemType/page/:index", api.GetItemPagedList)
	// r.GET("/:signableAccount/list/:itemType", api.GetItemList)

	// // links
	// r.GET("/:signableAccount/list/links/:linkUri/:index", api.GetLinkList)
	// r.GET("/:signableAccount/list/backlinks/:linkUri", api.GetBacklinkList)

	return r
}
