#!/bin/bash

CGO_ENABLED=0 GOOS=linux go build -o hello.exe -a -tags netgo -ldflags -w hello.go

docker build -t gcr.io/foolproof-io/kv-store:v1.0 .

rm hello.exe
