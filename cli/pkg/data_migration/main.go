package data_migration

import "flag"

var (
	fromDirectory string
)

func Start() {
	flag.StringVar(&fromDirectory, "from", "", "Directory to read v0.3.1 files from")
	flag.Parse()
	if fromDirectory == "" {
		flag.PrintDefaults()
		return
	}
	migrate(fromDirectory)
}
