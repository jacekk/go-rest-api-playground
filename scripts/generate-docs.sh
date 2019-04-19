#!/usr/bin/env bash

### Based on https://github.com/swaggo/gin-swagger#start-using-it

ROOT_DIR=`dirname $0`
DOCS_DIR="$ROOT_DIR/../docs"
CHECK_FILE="$DOCS_DIR/docs.go"
FILES_TO_REMOVE="$DOCS_DIR/*.*"

echo $ROOT_DIR

if [ -f $CHECK_FILE ]; then
    rm -rf $FILES_TO_REMOVE
fi

if ! hash swag 2>/dev/null; then
    go get -u github.com/swaggo/swag/cmd/swag
fi

swag init
