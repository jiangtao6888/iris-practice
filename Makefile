#!/bin/bash
export LANG=zh_CN.UTF-8

ENVARG=GOPATH=$(CURDIR)
LINUXARG=CGO_ENABLED=1 GOOS=linux GOARCH=amd64
BUILDARG=-ldflags " -s -X main.buildTime=`date '+%Y-%m-%dT%H:%M:%S'` -X main.gitHash=(`git symbolic-ref --short -q HEAD`)`git rev-parse HEAD`"

export GOPATH

dep:
	cd src; ${ENVARG} go get ./...; cd -

p:
	mkdir -p src/libraries/proto
	rm -rf src/libraries/proto/*
	cd src; protoc -I ../protocol --gofast_out=plugins=grpc:. common.proto; cd -
	cd src; protoc -I ../protocol --gofast_out=plugins=grpc:. user.proto; cd -

	ls src/libraries/proto/*.pb.go | xargs sed -i -e "s/,omitempty//"
	ls src/libraries/proto/*.pb.go | xargs sed -i -e "s@\"libraries/proto/@\"iris/libraries/proto/@"
	rm -f src/libraries/proto/*.pb.go-e
iris:
	cd src; ${ENVARG} go build ${BUILDARG} -o ../bin/iris main.go;
