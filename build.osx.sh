#!/bin/bash
cd ui
yarn bundle
cd ..
go-bindata -prefix "ui/build" -o assets.go ui/build/index.dist.js
go build -gccgoflags "-lusb-1.0" -tags dist -o dist/osx/Wally.app/Contents/MacOS/Wally
go build -gccgoflags "-lusb-1.0" -o dist/osx/wally-cli cli/main.go
upx dist/osx/wally-cli
