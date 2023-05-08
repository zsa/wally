#!/bin/sh

target=11
aldsfj;adlsjf

export CGO_LDFLAGS=-mmacosx-version-min=$target
export CGO_CFLAGS=-mmacosx-version-min=$target
export CGO_CXXFLAGS=-mmacosx-version-min=$target

echo "OSX Target version set to v${target}"