package data_migration

import (
	"flag"
	"log"
)

var (
	fromDirectory string
	deleteFile    bool
)

func init() {
	flag.StringVar(&fromDirectory, "from", "", "Directory to read v0.3.1 files from")
	flag.BoolVar(&deleteFile, "delete", false, "Delete files after migration")
}

func Start() {
	flag.Parse()
	if fromDirectory == "" {
		flag.PrintDefaults()
		return
	}
	if err := prepareDB(); err != nil {
		log.Fatalln(err)
	}
	if err := migrate(fromDirectory, deleteFile); err != nil {
		log.Fatalln(err)
	}
}
