#!/usr/bin/env bash

service=${1}

# check variables
if [ -z ${service} ]
then
    echo "First argument (service) is missing"
    exit 1
fi

echo "Running: protoc -I ${GOPATH}/src ${GOPATH}/src/bitbucket.org/pandaoj/go/${service}/pb/${service}.proto --go_out=plugins=grpc:${GOPATH}/src"

protoc -I ${GOPATH}/src ${GOPATH}/src/bitbucket.org/pandaoj/go/${service}/pb/${service}.proto --go_out=plugins=grpc:${GOPATH}/src