#!/usr/bin/env bash

go get -u github.com/golang/lint/golint

res=$(golint ./...)
if [ -n "$res" ]
then
    golint ./...;
    exit 1
fi
echo "golint passed"
exit 0