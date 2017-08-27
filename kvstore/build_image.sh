#!/bin/bash

GOPATH=$(pwd)

CGO_ENABLED=0 GOOS=linux go build -o server.exe -a -tags netgo -ldflags -w foolproof.io/kvstore/server

docker build -t gcr.io/foolproof-io/kvstore:v1.1 .
