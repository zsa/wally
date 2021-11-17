#!/bin/bash

SCRIPT_DIR="$(builtin cd "$(dirname "${BASH_SOURCE[0]}")" && pwd -P)"

cd "$SCRIPT_DIR"

CGO_CFLAGS=-mmacosx-version-min=10.8 \
CGO_LDFLAGS=-mmacosx-version-min=10.8 \
wails build

mkdir -p ./dist/osx/Wally.app/Contents/MacOS
mkdir -p ./dist/osx/Wally.app/Contents/libs

mv ./build/wally ./dist/osx/Wally.app/Contents/MacOS/Wally
dylibbundler -of -b -x ./dist/osx/Wally.app/Contents/MacOS/Wally -d ./dist/osx/Wally.app/Contents/libs/
