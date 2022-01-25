#!/bin/sh

go build -o dist/GI hub/hub.go
go build -o dist/RN indexer/indexer.go
