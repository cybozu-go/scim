#!/bin/bash

DIR=$(cd $(dirname $0); pwd -P)

set -e
sketch \
	--tmpl-dir=$DIR/tmpl \
	--with-builders \
	--with-has-methods \
	--with-key-prefix \
	-d $(cd $DIR/../resource; pwd -P) \
	$DIR/schema 
