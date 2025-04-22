#!/bin/sh
set -e

if [ "bundle" = "$1" ]; then
  bundle install
fi

exec "$@"
