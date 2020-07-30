#!/bin/bash

wails build -ldflags "-X main.mmacosx-version-min=10.8"
mv ./build/wally ./dist/osx/Wally.app/Contents/MacOS/Wally
cd ./dist/osx/Wally.app/Contents/MacOS/
dylibbundler -x Wally
