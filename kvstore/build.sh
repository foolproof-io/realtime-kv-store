#!/bin/bash

GOPATH=$(pwd)

CGO_ENABLED=0 GOOS=linux go build -o kvstore.exe -a -tags netgo -ldflags -w foolproof.io/kvstore

docker build -t gcr.io/foolproof-io/kvstore:v1.0 .
