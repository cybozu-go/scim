#!/bin/bash

# Script to perform code generation. This exists to overcome
# the fact that go:generate doesn't really allow you to change directories

# This file is expected to be executed from schema directory

set -e

echo "👉 Generating schema..."
DIR=../tools/cmd/genschema

pushd "$DIR" > /dev/null
go build -o .genschema main.go
popd > /dev/null

EXE="${DIR}/.genschema"
"$EXE" -schema=$DIR/schema.yml
echo "✔ done!"

rm "$EXE"
