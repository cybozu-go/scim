#!/bin/bash

# Script to perform code generation. This exists to overcome
# the fact that go:generate doesn't really allow you to change directories

# This file is expected to be executed from filter directory

set -e

echo "👉 Generating filter..."
DIR=../tools/cmd/genfilter

pushd "$DIR" > /dev/null
go build -o .genfilter main.go
popd > /dev/null

EXE="${DIR}/.genfilter"
"$EXE" -objects=$DIR/objects.yml
echo "✔ done!"

rm "$EXE"
