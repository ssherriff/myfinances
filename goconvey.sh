#!/bin/sh

go get github.com/smartystreets/goconvey

# since it's possible to set multiple gopaths, let's get first path
GP=${GOPATH%%:*}

$GP/bin/goconvey