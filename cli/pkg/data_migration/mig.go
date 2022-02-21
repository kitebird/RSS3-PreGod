package data_migration

import (
	"os"
	"path/filepath"
	"strings"
)

func migrate(fromDir string) error {
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
		if !strings.Contains(filename, "-") {
			// is main index
			return handleMainIndex(filename)
		} else if strings.Contains(filename, "-list-assets") {
			return handleAutoAssetList(filename)
		} else if strings.Contains(filename, "-list-links") {
			return handleLinkList(filename)
		} else if strings.Contains(filename, "-list-backlinks") {
			return handleLinkBackList(filename)
		} else {
			panic("unknown file: " + filename)
		}

	})

}

func handleMainIndex(file string) error {
	// handle main index

	return nil
}

func handleLinkList(file string) error {
	// handle link list

	return nil
}

func handleLinkBackList(file string) error {
	// handle link back list

	return nil
}

func handleAutoAssetList(file string) error {
	// handle auto asset list

	return nil
}
