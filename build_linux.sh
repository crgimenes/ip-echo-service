#!/bin/bash

#GOOS=linux GOARCH=amd64 go build
GOARCH=amd64 CGO_ENABLED=0 GOOS=linux \
    go build -a -ldflags '-s -w' -o ip-echo-service


