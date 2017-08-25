#!/bin/bash

CGO_ENABLED=0 GOOS=linux go build -o hello.exe -a -tags netgo -ldflags -w hello.go

docker build -t hello .

rm hello.exe
