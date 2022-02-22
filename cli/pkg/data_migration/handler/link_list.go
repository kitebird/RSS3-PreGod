package handler

import (
	"encoding/json"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/cli/pkg/data_migration/protocol"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/hub/pkg/db"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/hub/pkg/db/model"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/constants"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

func getLinkType(linkType string) constants.LinkTypeID {
	switch strings.ToLower(linkType) {
	case "following":
		return constants.LinkType_Following
	case "comment":
		return constants.LinkType_Comment
	case "like":
		return constants.LinkType_Like
	case "collection":
		return constants.LinkType_Collection
	default:
		return constants.LinkType_Unknown
	}
}

func LinkList(filebytes []byte) error {
	// handle link list
	var linkList protocol.RSS3Links031
	// Unmarshal
	if err := json.Unmarshal(filebytes, &linkList); err != nil {
		return err
	}
	// Split & save into db

	splits := strings.Split(linkList.ID, "-") // {account} - list - links.{id} - {pageNumber}

	account := splits[0]
	listID := strings.Split(splits[2], ".")[1]
	pageNumberStr := splits[3]

	instanceID := "rss3://account:" + account + "@evm"
	linkListID := instanceID + "/list/links/" + listID + "/" + pageNumberStr

	pageNumber, err := strconv.Atoi(pageNumberStr)
	if err != nil {
		return err
	}

	// Generate link list
	var links []model.Link

	for _, link := range linkList.List {
		targetInstanceID := "rss3://account:" + link + "@evm"
		links = append(links, model.Link{
			LinkID: linkListID, // todo: what is this ? uuid ?
			//ItemID:       "",
			RSS3ID:       instanceID,
			Prefix:       constants.Prefix_Account,
			TargetRSS3ID: targetInstanceID,
			TargetPrefix: constants.Prefix_Account,
			PageIndex:    pageNumber,
		})
	}

	// todo: create link_metadata ?

	return db.DB.Transaction(func(tx *gorm.DB) error {
		return tx.Create(&links).Error
	})
}
