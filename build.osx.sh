#!/bin/bash

CGO_CFLAGS=-mmacosx-version-min=10.8 \
CGO_LDFLAGS=-mmacosx-version-min=10.8 \
wails build
mv ./build/wally ./dist/osx/Wally.app/Contents/MacOS/Wally
cd ./dist/osx/Wally.app/Contents/MacOS/
dylibbundler -x Wally
