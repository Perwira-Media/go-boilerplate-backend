#!/bin/sh
echo ">>> Compiling Boiler Code into Alpine's binary file..."
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./_bin/pm-boiler-code -v ./cmd/boiler-plate/main.go
if [ $? -ne 0 ]; then
    echo "->>> Compiling Failed."
    exit -1
fi
echo "->>> Compiling Success."
echo "->>> Buidling into docker image..."
version=`cat VERSION`
if [ "$1" = "dev" ]; then
    version=$version"_dev"
fi
docker build -t perwiramedia/pm-boiler-code:latest -t perwiramedia/pm-boiler-code:$version -f Dockerfile .
if [ $? -ne 0 ]; then
    echo "->>> Building Failed."
    exit -1
fi
echo "->>> Building Success."