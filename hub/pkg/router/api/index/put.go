package index

import (
	"github.com/gin-gonic/gin"
)

type PutIndexRequest struct{}

func PutIndexHandlerFunc(c *gin.Context) {

	// Parse the request

	// value, exists := c.Get(middleware.KeyInstance)
	// if !exists {
	// 	w := web.Gin{C: c}
	// 	w.JSONResponse(http.StatusBadRequest, status.CodeInvalidParams, nil)

	// 	return
	// }

	// platformInstance, ok := value.(*rss3uri.PlatformInstance)
	// if !ok {
	// 	w := web.Gin{C: c}
	// 	w.JSONResponse(http.StatusBadRequest, status.CodeInvalidParams, nil)

	// 	return
	// }

	// var indexFile protocol.Index

	// if err := c.ShouldBind(&indexFile); err != nil {
	// 	w := web.Gin{C: c}
	// 	w.JSONResponse(http.StatusBadRequest, status.CodeInvalidParams, nil)

	// 	return
	// }

}
