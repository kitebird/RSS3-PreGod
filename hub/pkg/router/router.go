package router

import (
	"net/http"

	"github.com/NaturalSelectionLabs/RSS3-PreGod/hub/pkg/middleware"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/hub/pkg/router/api"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/hub/pkg/router/monitor"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	// Apply middlewares
	r.Use(middleware.Logger(), gin.Recovery())

	// === APIs ===

	// instance
	r.GET("/:authority", api.GetInstance)

	// // items
	// r.GET("/:authority/:item_type/:item_uuid", api.GetItem)
	// r.GET("/:authority/list/:item_type/:page_index", api.GetItemPagedList)
	// r.GET("/:authority/list/:item_type", api.GetItemList)

	// // links
	// r.GET("/:authority/list/links/:link_type/:page_index", api.GetLinkList)
	// r.GET("/:authority/list/backlinks/:link_type", api.GetBacklinkList)

	// === Monitor ===
	r.GET("/ping", api.Ping)
	r.GET("/debug/statsviz/*filepath", monitor.Statsviz)

	// === Static ===
	r.GET("/favicon.ico", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://rss3.io/favicon.ico")
	})

	return r
}
