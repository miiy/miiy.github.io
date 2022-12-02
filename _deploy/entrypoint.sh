#!/bin/sh
set -e

if [ "jekyll" == "$1" ]; then
  bundle install
fi

exec "$@"
