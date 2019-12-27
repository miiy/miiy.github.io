#!/bin/bash
set -e

if [ ! -f Gemfile ]; then
  echo "Error: Cant't find Gemfile"
  exit 1
fi

bundle install

exec "$@"
