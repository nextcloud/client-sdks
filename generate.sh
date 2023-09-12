#!/bin/sh
set -euxo

cd "$(dirname "$0")"

docker build -t nextcloud-client-sdks -f Dockerfile docker/

docker run --rm -v "$(pwd):/src" nextcloud-client-sdks
