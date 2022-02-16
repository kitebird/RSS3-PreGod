package router

import (
	"github.com/NaturalSelectionLabs/RSS3-PreGod/hub/pkg/router/api"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/hub/pkg/router/monitor"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.GET("/ping", api.Ping)

	// instance
	r.GET("/:authority", api.GetInstance)

	// // items
	// r.GET("/:authority/:item_type/:item_uuid", api.GetItem)
	// r.GET("/:authority/list/:item_type/:page_index", api.GetItemPagedList)
	// r.GET("/:authority/list/:item_type", api.GetItemList)

	// // links
	// r.GET("/:authority/list/links/:link_type/:page_index", api.GetLinkList)
	// r.GET("/:authority/list/backlinks/:link_type", api.GetBacklinkList)

	// monitor
	r.GET("/debug/statsviz/*filepath", monitor.Statsviz)

	return r
}
