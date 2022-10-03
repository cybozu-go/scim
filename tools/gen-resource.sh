#!/bin/bash

DIR=$(cd $(dirname $0); pwd -P)

set -e
sketch \
	--dev-mode --dev-path=$DIR/../../../lestrrat-go/sketch \
	--verbose \
	--tmpl-dir=$DIR/tmpl/resource \
	--exclude='decodeExtraField$' \
	--with-key-name-prefix \
	-d $(cd $DIR/../resource; pwd -P) \
	$DIR/schema 
