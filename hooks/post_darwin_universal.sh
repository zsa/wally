#!/bin/sh

rm ./wally.dmg >/dev/null 2>&1
rm ./wally.app/Contents/libs >/dev/null 2>&1
mkdir ./wally.app/Contents/libs >/dev/null 2>&1
dylibbundler -of -b -x ./wally.app/Contents/MacOS/wally -d ./wally.app/Contents/libs -s /opt/homebrew/Cellar/libusb/1.0.26/lib
appdmg ../darwin/dmg.json ./wally.dmg