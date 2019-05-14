#!/bin/bash
cd ui
yarn bundle
cd ..
go-bindata -prefix "ui/build" -o assets.go ui/build/index.dist.js
go build -tags dist -o dist/linux64/wally
go build -o dist/linux64/wally-cli cli/main.go
upx dist/linux64/wally
upx dist/linux64/wally-cli
