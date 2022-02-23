package handler

import (
	"encoding/json"
	"errors"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/cli/pkg/data_migration/protocol"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/hub/pkg/db"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/hub/pkg/db/model"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/constants"
	"gorm.io/gorm"
	"strings"
)

func getAccountPlatform(platform string) constants.PlatformNameID {
	switch strings.ToLower(platform) {
	case "evm+":
		return constants.PlatformNameID_Evm
	case "solana":
		return constants.PlatformNameID_Solana
	case "flow":
		return constants.PlatformNameID_Flow
	case "twitter":
		return constants.PlatformNameID_Twitter
	case "misskey":
		return constants.PlatformNameID_Misskey
	case "jike":
		return constants.PlatformNameID_Jike
	default:
		return constants.PlatformNameID_Unknown
	}
}

func MainIndex(filebytes []byte) error {
	// handle main index
	var mainIndex protocol.RSS3Index031
	// Unmarshal
	if err := json.Unmarshal(filebytes, &mainIndex); err != nil {
		return err
	}
	// Split & save into db

	var instanceBase model.InstanceBase

	newID := mainIndex.ID + "@" + string(constants.PlatformName_Evm)

	if err := db.DB.First(&instanceBase, "RSS3ID = ?", newID).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		// Already exists
		return nil // skip
	}

	// New instance
	instanceBase = model.InstanceBase{
		RSS3ID:         newID,
		Prefix:         constants.Prefix_Account,
		InstanceTypeID: constants.InstanceType_Account,
	}

	// Accounts
	accounts := []model.AccountPlatform{
		{
			AccountID:         newID,
			PlatformNameID:    constants.PlatformNameID_Evm,
			PlatformAccountID: mainIndex.ID,
		},
	}

	for _, additionalAccount := range mainIndex.Profile.Accounts {
		splits := strings.Split(additionalAccount.ID, "-")
		accounts = append(accounts, model.AccountPlatform{
			AccountID:         splits[0] + "@" + splits[1],
			PlatformNameID:    getAccountPlatform(splits[0]),
			PlatformAccountID: splits[1],
		})
	}

	// Account
	account := &model.Account{
		AccountID: newID,
		Name:      mainIndex.Profile.Name,
		Bio:       mainIndex.Profile.Bio,
		Avatars:   mainIndex.Profile.Avatar,

		InstanceBase:    instanceBase,
		AccountPlatform: accounts,
	}

	// todo: migrate application data (signature, account tags, custom fields, etc)

	// Save
	return db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&instanceBase).Error; err != nil {
			return err
		}
		if err := tx.Save(&accounts).Error; err != nil {
			return err
		}
		if err := tx.Create(&account).Error; err != nil {
			return err
		}
		return nil
	})
}
