#!/bin/bash

DIR=$(cd $(dirname $0); pwd -P)

set -e
sketch \
	--tmpl-dir=$DIR/tmpl/resource \
	--with-builders \
	--with-has-methods \
	--with-key-name-prefix \
	-d $(cd $DIR/../resource; pwd -P) \
	$DIR/schema 
