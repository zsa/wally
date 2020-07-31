#!/bin/bash

CGO_CFLAGS=-mmacosx-version-min=10.8 \
CGO_LDFLAGS=-mmacosx-version-min=10.8 \
wails build
mv ./build/wally ./dist/osx/Wally.app/Contents/MacOS/Wally
dylibbundler -of -b -x ./dist/osx/Wally.app/Contents/MacOS/Wally -d ./dist/osx/Wally.app/Contents/libs/
