#!/bin/sh

mkdir ./wally.app/Contents/libs
dylibbundler -of -b -x ./wally.app/Contents/MacOS/wally -d ./wally.app/Contents/libs -s /usr/local/lib
appdmg ../darwin/dmg.json ./wally.dmg