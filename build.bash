#!/bin/bash

VERSION=`git describe --tags`
go build -ldflags "-X main.version=${VERSION}"
echo "build as ${VERSION}"
