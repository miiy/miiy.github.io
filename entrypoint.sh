#!/bin/sh
set -e

# Dependencies are installed into the image at build time (see Dockerfile),
# so there is nothing to install at container start.
exec "$@"
