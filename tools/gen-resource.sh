#!/bin/bash

PATH="$(pwd -P)/../../lestrrat-go/sketch:$PATH"

DIR=$(cd $(dirname $0); pwd -P)

set -e
sketch \
	--dev-mode --dev-path=$(pwd -P)/../../lestrrat-go/sketch \
	--tmpl-dir=$DIR/tmpl/resource \
	--with-builders \
	--with-has-methods \
	--with-key-name-prefix \
	-d $(cd $DIR/../resource; pwd -P) \
	$DIR/schema 
