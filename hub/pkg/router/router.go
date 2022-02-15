package router

import (
	"github.com/NaturalSelectionLabs/RSS3-PreGod/hub/pkg/router/api"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.GET("/ping", api.Ping)

	// instance
	r.GET("/:instance", api.GetInstance)

	// // items
	// r.GET("/:instance/:item_type/:item_uuid", api.GetItem)
	// r.GET("/:instance/list/:item_type/:page_index", api.GetItemPagedList)
	// r.GET("/:instance/list/:item_type", api.GetItemList)

	// // links
	// r.GET("/:instance/list/links/:link_type/:page_index", api.GetLinkList)
	// r.GET("/:instance/list/backlinks/:link_type", api.GetBacklinkList)

	return r
}
