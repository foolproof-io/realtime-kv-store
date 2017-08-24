#!/bin/bash

CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags -w hello.go
docker build -t hello .
