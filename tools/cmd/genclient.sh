#!/bin/bash

# Script to perform code generation. This exists to overcome
# the fact that go:generate doesn't really allow you to change directories

# This file is expected to be executed from resource directory

set -e

echo "👉 Generating client..."
DIR=../tools/cmd/genclient

pushd "$DIR" > /dev/null
GOWORK=off go build -o .genclient main.go
popd > /dev/null

EXE="${DIR}/.genclient"
"$EXE" -calls=$DIR/calls.yml
echo "✔ done!"

rm "$EXE"
