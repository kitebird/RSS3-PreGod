package index

import (
	"context"
	"net/http"

	"github.com/NaturalSelectionLabs/RSS3-PreGod/hub/database"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/hub/database/model"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/hub/pkg/middleware"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/hub/pkg/protocol"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/hub/pkg/status"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/hub/pkg/web"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/rss3uri"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PutIndexRequest struct{}

func PutIndexHandlerFunc(c *gin.Context) {
	w := web.Gin{C: c}

	// Parse the request
	value, exists := c.Get(middleware.KeyInstance)
	if !exists {
		w.JSONResponse(http.StatusBadRequest, status.CodeInvalidParams, nil)

		return
	}

	platformInstance, ok := value.(*rss3uri.PlatformInstance)
	if !ok {
		w.JSONResponse(http.StatusBadRequest, status.CodeInvalidParams, nil)

		return
	}

	// Get the new put index file
	var indexFile protocol.Index
	if err := c.ShouldBind(&indexFile); err != nil {
		w.JSONResponse(http.StatusBadRequest, status.CodeInvalidParams, nil)

		return
	}

	// TODO: Check the signature

	// Get the original index file
	oIndexFile, httpStatus, bizStatus := getIndexFile(platformInstance)
	if oIndexFile == nil {
		w.JSONResponse(httpStatus, bizStatus, nil)

		return
	}

	// Start the transaction
	tx := database.Instance.Tx(context.Background())
	defer tx.Rollback()

	// Compare the diff
	httpStatus, bizStatus = updateAccount(tx, &indexFile, oIndexFile)
	if httpStatus != http.StatusOK {
		w.JSONResponse(httpStatus, bizStatus, nil)

		return
	}

	// TODO: compare more diffs

	return //nolint:gosimple // TODO:
}

func updateAccount(db *gorm.DB, indexFile *protocol.Index, oIndexFile *protocol.Index) (int, status.Code) {
	toUpdate := model.Account{}

	if indexFile.Profile.Name != oIndexFile.Profile.Name {
		toUpdate.Name = database.WrapNullString(*indexFile.Profile.Name)
	}

	if indexFile.Profile.Bio != oIndexFile.Profile.Bio {
		toUpdate.Bio = database.WrapNullString(*indexFile.Profile.Bio)
	}

	if !isArrayEqual(indexFile.Profile.Avatars, oIndexFile.Profile.Avatars) {
		toUpdate.Avatars = indexFile.Profile.Avatars
	}

	if !isAttachmentsEqual(indexFile.Profile.Attachments, oIndexFile.Profile.Attachments) {
		toUpdate.Attachments = indexFile.Profile.Attachments.ToDBStruct()
	}

	if err := db.Updates(&toUpdate).Error; err != nil {
		return http.StatusInternalServerError, status.CodeError
	}

	return http.StatusOK, status.CodeSuccess
}

func isArrayEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func isAttachmentsEqual(a, b []protocol.IndexProfileAttachment) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		aa := a[i]
		bb := b[i]

		if aa.Type != bb.Type ||
			aa.Content != bb.Content ||
			aa.Address != bb.Address ||
			aa.MimeType != bb.MimeType ||
			aa.SizeInBytes != bb.SizeInBytes {
			return false
		}
	}

	return true
}
