package data_migration

import "flag"

var (
	fromDirectory string
)

func init() {
	flag.StringVar(&fromDirectory, "from", "", "Directory to read v0.3.1 files from")
}

func Start() {
	flag.Parse()
	if fromDirectory == "" {
		flag.PrintDefaults()
		return
	}
	migrate(fromDirectory)
}
