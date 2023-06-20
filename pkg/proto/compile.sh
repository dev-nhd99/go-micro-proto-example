#!/bin/bash
BASEDIR=$(dirname $0)
echo "Script location: ${BASEDIR}"
#protoc ./$BASEDIR/calculator.proto --micro_out=. --go_out=. ./$BASEDIR/calculator.proto
protoc --go_out=. --go-grpc_out=. --micro_out=. ./calculator.proto