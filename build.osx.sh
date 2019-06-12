#!/bin/bash
cd ui
yarn bundle
cd ..
go-bindata -prefix "ui/build" -o assets.go ui/build/index.dist.js
CGO_LDFLAGS="-mmacosx-version-min=10.8 -lusb-1.0" \
CGO_CFLAGS=-mmacosx-version-min=10.8 \
go build -tags dist -o dist/osx/Wally.app/Contents/MacOS/Wally
CGO_LDFLAGS="-mmacosx-version-min=10.8 -lusb-1.0" \
CGO_CFLAGS=-mmacosx-version-min=10.8 \
go build -o dist/osx/wally-cli cli/main.go
upx dist/osx/wally-cli
