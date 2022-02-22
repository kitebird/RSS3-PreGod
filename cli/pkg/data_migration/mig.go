package data_migration

import (
	"encoding/json"
	"errors"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/cli/pkg/data_migration/model"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func migrate(fromDir string, delete bool) error {

	// what file:
	// 1. main index file
	// 2. link list (following)
	// 3. link backlist (following)
	// 4. auto asset list

	// for each filename, use specific handler
	return filepath.Walk(fromDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			// skip
			return nil
		}
		// do something
		filename := info.Name()
		filebytes, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		var errInHandle error = nil
		if !strings.Contains(filename, "-") {
			// is main index
			errInHandle = handleMainIndex(filebytes)
		} else if strings.Contains(filename, "-list-assets") {
			errInHandle = handleAutoAssetList(filebytes)
		} else if strings.Contains(filename, "-list-links") {
			errInHandle = handleLinkList(filebytes)
		} else if strings.Contains(filename, "-list-backlinks") {
			errInHandle = handleLinkBackList(filebytes)
		} else {
			errInHandle = errors.New("Unknown file: " + filename)
		}

		if errInHandle == nil && delete {
			// delete
			errInHandle = os.Remove(path)
		}

		log.Println("Success: ", filename)

		return errInHandle

	})

}

func handleMainIndex(filebytes []byte) error {
	// handle main index
	var mainIndex model.RSS3Index031
	// Unmarshal
	if err := json.Unmarshal(filebytes, &mainIndex); err != nil {
		return err
	}
	// Split & save into db

	return nil
}

func handleLinkList(filebytes []byte) error {
	// handle link list
	var linkList model.RSS3Links031
	// Unmarshal
	if err := json.Unmarshal(filebytes, &linkList); err != nil {
		return err
	}
	// Split & save into db

	return nil
}

func handleLinkBackList(filebytes []byte) error {
	// handle link back list
	var linkBackList model.RSS3Backlinks031
	// Unmarshal
	if err := json.Unmarshal(filebytes, &linkBackList); err != nil {
		return err
	}
	// Split & save into db

	return nil
}

func handleAutoAssetList(filebytes []byte) error {
	// handle auto asset list
	var autoAssetList model.RSS3AutoAssets031
	// Unmarshal
	if err := json.Unmarshal(filebytes, &autoAssetList); err != nil {
		return err
	}
	// Split & save into db

	return nil
}
