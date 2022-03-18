package index

import (
	"context"
	"net/http"
	"sync"

	"github.com/NaturalSelectionLabs/RSS3-PreGod/hub/database"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/hub/database/model"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/hub/pkg/middleware"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/hub/pkg/protocol"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/hub/pkg/status"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/hub/pkg/web"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/constants"
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
	tx := database.Instance.DB(context.Background()).Begin()

	wg := sync.WaitGroup{}

	const MaxTask = 10

	ch := make(chan []int, MaxTask)
	defer close(ch)

	wg.Add(1)

	go func() {
		httpStatus, bizStatus = runUpdateTask(updateAccount, tx, &indexFile, oIndexFile)

		wg.Done()
	}()

	wg.Add(1)

	go func() {
		httpStatus, bizStatus = runUpdateTask(updateAccountPlatform, tx, &indexFile, oIndexFile)

		wg.Done()
	}()

	wg.Wait()

	if err := tx.Commit().Error; err != nil {
		w.JSONResponse(http.StatusInternalServerError, status.CodeError, nil)

		return
	}

	w.JSONResponse(http.StatusOK, status.CodeSuccess, indexFile)

	// TODO: compare more diffs

	return //nolint:gosimple // TODO:
}

type UpdateTaskFunc func(db *gorm.DB, old, newIndex *protocol.Index) (int, status.Code)

func runUpdateTask(f UpdateTaskFunc, db *gorm.DB, oldIndex, newIndex *protocol.Index) (int, status.Code) {
	httpStatus, bizStatus := f(db, oldIndex, newIndex)

	return httpStatus, bizStatus
}

func updateAccount(db *gorm.DB, indexFile *protocol.Index, oIndexFile *protocol.Index) (int, status.Code) {
	u, err := rss3uri.Parse(indexFile.Identifier)
	if err != nil {
		return http.StatusInternalServerError, status.CodeError
	}

	needUpdate := false

	toUpdate := model.Account{
		ID: u.Instance.GetIdentity(),
	}

	if *indexFile.Profile.Name != *oIndexFile.Profile.Name {
		needUpdate = true
		toUpdate.Name = database.WrapNullString(*indexFile.Profile.Name)
	}

	if *indexFile.Profile.Bio != *oIndexFile.Profile.Bio {
		needUpdate = true
		toUpdate.Bio = database.WrapNullString(*indexFile.Profile.Bio)
	}

	if !isArrayEqual(indexFile.Profile.Avatars, oIndexFile.Profile.Avatars) {
		needUpdate = true
		toUpdate.Avatars = indexFile.Profile.Avatars
	}

	if !isAttachmentsEqual(indexFile.Profile.Attachments, oIndexFile.Profile.Attachments) {
		needUpdate = true
		toUpdate.Attachments = indexFile.Profile.Attachments.ToDBStruct()
	}

	if needUpdate {
		if err := db.Updates(&toUpdate).Error; err != nil {
			return http.StatusInternalServerError, status.CodeError
		}
	}

	return http.StatusOK, status.CodeSuccess
}

// nolint:funlen,gocognit // TODO
func updateAccountPlatform(db *gorm.DB, indexFile *protocol.Index, oIndexFile *protocol.Index) (int, status.Code) {
	var toUpdate, toCreate, toDelete []model.AccountPlatform

	oldAccountMap := make(map[string]string)
	newAccountMap := make(map[string]string)

	for _, account := range indexFile.Profile.Accounts {
		newAccountMap[account.Identifier] = account.Signature
	}

	for _, account := range oIndexFile.Profile.Accounts {
		oldAccountMap[account.Identifier] = account.Signature
	}

	accountURI, err := rss3uri.Parse(indexFile.Identifier)
	if err != nil {
		return http.StatusInternalServerError, status.CodeError
	}

	for _, account := range indexFile.Profile.Accounts {
		signature, exist := oldAccountMap[account.Identifier]
		if !exist {
			platformURI, err := rss3uri.Parse(account.Identifier)
			if err != nil {
				return http.StatusInternalServerError, status.CodeError
			}

			toCreate = append(toCreate, model.AccountPlatform{
				AccountID:         accountURI.Instance.GetIdentity(),
				AccountPlatformID: int(constants.PlatformSymbol(accountURI.Instance.GetSuffix()).ID()),
				PlatformAccountID: platformURI.Instance.GetIdentity(),
				PlatformID:        int(constants.PlatformSymbol(platformURI.Instance.GetSuffix()).ID()),
			})

			continue
		}

		platformURI, err := rss3uri.Parse(account.Identifier)
		if err != nil {
			return http.StatusInternalServerError, status.CodeError
		}

		if signature != account.Signature {
			toUpdate = append(toUpdate, model.AccountPlatform{
				AccountID:         accountURI.Instance.GetIdentity(),
				AccountPlatformID: int(constants.PlatformSymbol(accountURI.Instance.GetSuffix()).ID()),
				PlatformAccountID: platformURI.Instance.GetIdentity(),
				PlatformID:        int(constants.PlatformSymbol(platformURI.Instance.GetSuffix()).ID()),
			})
		}
	}

	for _, account := range oIndexFile.Profile.Accounts {
		_, exist := newAccountMap[account.Identifier]
		if !exist {
			platformURI, err := rss3uri.Parse(account.Identifier)
			if err != nil {
				return http.StatusInternalServerError, status.CodeError
			}

			toDelete = append(toDelete, model.AccountPlatform{
				AccountID:         accountURI.Instance.GetIdentity(),
				AccountPlatformID: int(constants.PlatformSymbol(accountURI.Instance.GetSuffix()).ID()),
				PlatformAccountID: platformURI.Instance.GetIdentity(),
				PlatformID:        int(constants.PlatformSymbol(platformURI.Instance.GetSuffix()).ID()),
			})
		}
	}

	if len(toDelete) > 0 {
		// TODO
		for i := 0; i < len(toDelete); i++ {
			if err := db.Where(&toDelete[i]).Delete(&model.AccountPlatform{}).Error; err != nil {
				return http.StatusInternalServerError, status.CodeError
			}
		}
	}

	if len(toUpdate) > 0 {
		if err := db.Updates(&toUpdate).Error; err != nil {
			return http.StatusInternalServerError, status.CodeError
		}
	}

	if len(toCreate) > 0 {
		if err := db.Create(&toCreate).Error; err != nil {
			return http.StatusInternalServerError, status.CodeError
		}
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
