#!/usr/bin/env bash

set -e

if [[ -d $1 ]]; then
  echo "$1 already exists"
  exit 255
fi

mkdir -p "$1"
cp daynext/daynext.go "$1/$1.go"
cp daynext/daynext_test.go "$1/$1_test.go"
sed -i "s/daynext/$1/g" "$1/$1.go"
sed -i "s/daynext/$1/g" "$1/$1_test.go"
