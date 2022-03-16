package index

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/NaturalSelectionLabs/RSS3-PreGod/hub/database"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/hub/pkg/middleware"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/hub/pkg/protocol"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/hub/pkg/status"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/hub/pkg/web"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/constants"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/rss3uri"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GetIndexRequest struct{}

func GetIndexHandlerFunc(c *gin.Context) {
	w := web.Gin{C: c}

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

	if platformInstance.Prefix != constants.PrefixNameAccount || platformInstance.Platform != constants.PlatformSymbolEthereum {
		w.JSONResponse(http.StatusBadRequest, status.CodeInvalidParams, nil)

		return
	}

	indexFile, httpStatus, bizStatus := GetIndexFile(platformInstance)

	// Query link list
	linkList, err := database.Instance.QueryLinkList(tx, 1, account.ID, int(constants.PrefixIDAccount), account.Platform)
	if err != nil {
		return
	}

	// TODO No test data available
	var (
		notePageIndex  int
		assetPageIndex int
	)

	identifier := rss3uri.New(platformInstance).String()

	var (
		name *string = nil
		bio  *string = nil
	)

	if account.Name.Valid {
		name = &account.Name.String
	}

	if account.Bio.Valid {
		bio = &account.Bio.String
	}

	indexFile := protocol.Index{
		SignedBase: protocol.SignedBase{
			Base: protocol.Base{
				Version:    protocol.Version,
				Identifier: identifier,
			},
		},
		Profile: protocol.IndexProfile{
			Name:    name,
			Avatars: account.Avatars,
			Bio:     bio,
			// TODO No data available
			// Attachments: nil,
		},
		Links: protocol.IndexLinks{
			Identifiers: []protocol.IndexLinkIdentifier{
				{
					Type:             "following",
					IdentifierCustom: fmt.Sprintf("%s/list/link/following/%d", identifier, linkList.MaxPageIndex),
					Identifier:       fmt.Sprintf("%s/list/link/following", identifier),
				},
			},
			IdentifierBack: fmt.Sprintf("%s/list/backlink", identifier),
		},
		Items: protocol.IndexItems{
			Notes: protocol.IndexItemsNotes{
				IdentifierCustom: fmt.Sprintf("%s/list/note/%d", identifier, 0),
				Identifier:       fmt.Sprintf("%s/list/note", identifier),
			},
			Assets: protocol.IndexItemsAssets{
				IdentifierCustom: fmt.Sprintf("%s/list/asset/%d", identifier, 0),
				Identifier:       fmt.Sprintf("%s/list/asset", identifier),
			},
		},
	}

	// Start the transaction
	tx := database.Instance.Tx(context.Background())
	defer tx.Rollback()

	// Query the account
	account, err := database.Instance.QueryAccount(
		tx,
		instance.GetIdentity(),
		int(constants.PlatformSymbol(instance.GetSuffix()).ID()),
	)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, http.StatusConflict, status.CodeError
		}

		return nil, http.StatusInternalServerError, status.CodeError
	}

	indexFile.Profile.Name = account.Name
	indexFile.Profile.Avatars = account.Avatars
	indexFile.Profile.Bio = account.Bio

	// Query the linklists
	// TODO: query linklist table for different types of links
	// TODO: put max page index into linklist table's metadata field
	followingMaxPageIndex, err := database.Instance.QueryLinkWithMaxPageIndex(tx, 1, account.ID, account.Platform)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, http.StatusInternalServerError, status.CodeError
	}

	indexFile.Links.Identifiers = append(indexFile.Links.Identifiers, protocol.IndexLinkIdentifier{
		Type:             "following",
		IdentifierCustom: fmt.Sprintf("%s/list/link/following/%d", identifier, followingMaxPageIndex),
		Identifier:       fmt.Sprintf("%s/list/link/following", identifier),
	})

	// Query the account platforms
	accountPlatforms, err := database.Instance.QueryAccountPlatforms(tx, account.ID, account.Platform)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, http.StatusInternalServerError, status.CodeError
	}

	for _, accountPlatform := range accountPlatforms {
		indexFile.Profile.Accounts = append(indexFile.Profile.Accounts, protocol.IndexAccount{
			Identifier: rss3uri.New(&rss3uri.PlatformInstance{
				Prefix:   constants.PrefixNameAccount,
				Identity: accountPlatform.PlatformAccountID,
				Platform: constants.PlatformID(accountPlatform.PlatformID).Symbol(),
			}).String(),
		})
	}

	// Query the signature
	signature, err := database.Instance.QuerySignature(
		tx,
		fmt.Sprintf("%s@%s", account.ID, constants.PlatformID(account.Platform).Symbol()),
	)
	if err != nil {
		return nil, http.StatusInternalServerError, status.CodeError
	}

	indexFile.Signature = signature.Signature
	indexFile.Base.DateCreated = signature.CreatedAt.Format(time.RFC3339)
	indexFile.Base.DateUpdated = signature.UpdatedAt.Format(time.RFC3339)

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return nil, http.StatusInternalServerError, status.CodeError
	}

	return &indexFile, http.StatusOK, status.CodeSuccess
}
