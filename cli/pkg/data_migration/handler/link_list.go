package handler

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/NaturalSelectionLabs/RSS3-PreGod/cli/pkg/data_migration/protocol"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/hub/pkg/db"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/hub/pkg/db/model"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/constants"
	"github.com/gofrs/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"gorm.io/gorm"
)

//nolint:funlen,gocognit // no need to split
func LinkList(content bson.D) error {
	// handle link list
	var linkList protocol.RSS3Links031
	// Unmarshal
	doc, err := bson.Marshal(content)
	if err != nil {
		return err
	}

	if err = bson.Unmarshal(doc, &linkList); err != nil {
		return err
	}

	CreatedAt, err := time.Parse("2006-01-02T15:04:05.000Z", linkList.DateCreated)
	if err != nil {
		return err
	}

	UpdatedAt, err := time.Parse("2006-01-02T15:04:05.000Z", linkList.DateUpdated)
	if err != nil {
		return err
	}

	// Split & save into db

	splits := strings.Split(linkList.ID, "-") // {account} - list - links.{id} - {pageNumber}

	account := splits[0]
	listID := strings.Split(splits[2], ".")[1]
	linkTypeId := constants.StringToLinkTypeID(listID)
	pageNumberStr := splits[3]

	instanceID := account + "@" + string(constants.PlatformSymbolEthereum)
	linkListID := instanceID + "/list/links/" + listID + "/" + pageNumberStr

	pageNumber, err := strconv.Atoi(pageNumberStr)
	if err != nil {
		return err
	}

	linkListModel := model.LinkList{
		LinkListID: linkListID,
		RSS3ID:     instanceID,
		LinkType:   linkTypeId,
		BaseModel: model.BaseModel{
			CreatedAt: CreatedAt,
			UpdatedAt: UpdatedAt,
		},
	}

	// Generate link list
	var links []model.Link = make([]model.Link, len(linkList.List))

	for i, target := range linkList.List {
		uniqueName :=
			fmt.Sprint(linkTypeId) +
				"-" + account + "-" + fmt.Sprint(constants.PrefixIDAccount) + "-" + fmt.Sprint(constants.PlatformIDEthereum) +
				"-" + target + "-" + fmt.Sprint(constants.PrefixIDAccount) + "-" + fmt.Sprint(constants.PlatformIDEthereum)

		linkID := uuid.NewV5(uuid.NamespaceOID, uniqueName).String()
		links[i] = model.Link{
			LinkID:           linkID,
			RSS3ID:           account,
			PrefixID:         constants.PrefixIDAccount,
			PlatformID:       constants.PlatformIDEthereum,
			TargetRSS3ID:     target,
			TargetPrefixID:   constants.PrefixIDAccount,
			TargetPlatformID: constants.PlatformIDEthereum,
			PageIndex:        pageNumber,
			BaseModel: model.BaseModel{
				CreatedAt: CreatedAt,
				UpdatedAt: UpdatedAt,
			},
		}
	}

	// deduplicate links array by linkID
	var dedupedLinks []model.Link

	for _, link := range links {
		var found bool = false

		for _, l := range dedupedLinks {
			if l.LinkID == link.LinkID {
				found = true

				break
			}
		}

		if !found {
			dedupedLinks = append(dedupedLinks, link)
		}
	}

	signature := &model.Signature{
		FileURI:   fmt.Sprintf("rss3://account:%s@%s/list/link/following/0", account, fmt.Sprint(constants.PlatformIDEthereum)),
		Signature: linkList.Signature,
		BaseModel: model.BaseModel{
			CreatedAt: CreatedAt,
			UpdatedAt: UpdatedAt,
		},
	}

	return db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&linkListModel).Error; err != nil {
			return err
		}

		if len(dedupedLinks) > 0 {
			// insert in batch
			chunkSize := 1000
			for i := 0; i < len(dedupedLinks); i += chunkSize {
				end := i + chunkSize
				if end > len(dedupedLinks) {
					end = len(dedupedLinks)
				}

				if err := tx.Create(dedupedLinks[i:end]).Error; err != nil {
					return err
				}
			}
		}

		if err := tx.Create(&signature).Error; err != nil {
			return err
		}

		return nil
	})
}
