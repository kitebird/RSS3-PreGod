package data_migration

import (
	"errors"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/cli/pkg/data_migration/handler"
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
			errInHandle = handler.MainIndex(filebytes)
		} else if strings.Contains(filename, "-list-assets") {
			errInHandle = handler.AssetListAuto(filebytes)
		} else if strings.Contains(filename, "-list-links") {
			errInHandle = handler.LinkList(filebytes)
		} else if strings.Contains(filename, "-list-backlinks") {
			errInHandle = handler.LinkBackList(filebytes)
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
