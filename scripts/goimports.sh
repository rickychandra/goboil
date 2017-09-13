#!/usr/bin/env bash

go get golang.org/x/tools/cmd/goimports

res=$(goimports -d .)
if [ -n "$res" ]
then
    goimports -d .;
    exit 1
fi
echo "goimports passed"
exit 0