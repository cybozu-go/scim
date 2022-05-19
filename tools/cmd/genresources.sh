#!/bin/bash

# Script to perform code generation. This exists to overcome
# the fact that go:generate doesn't really allow you to change directories

# This file is expected to be executed from resource directory

set -e

echo "👉 Generating resources..."
DIR=../tools/cmd/genresources

pushd "$DIR" > /dev/null
go build -o .genresources main.go
popd > /dev/null

EXE="${DIR}/.genresources"
"$EXE" -objects=$DIR/objects.yml
echo "✔ done!"

rm "$EXE"
