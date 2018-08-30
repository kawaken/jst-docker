#!/bin/bash

GOOS=linux GOARCH=amd64 CGO_ENABLE=0 go build -o time-test

for i in 1 2 3; do
  docker build -t jst-time-test$i -f Dockerfile$i .
done
