@echo off

go build -o dist/GI.exe hub/hub.go
go build -o dist/RN.exe indexer/indexer.go
