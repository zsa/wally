#!/usr/bin/env bash

set -eou pipefail

CONTAINER_NAME="$(uuidgen)"

docker build -t wally .
docker run --rm -d --name "$CONTAINER_NAME" wally

docker cp $CONTAINER_NAME:/project/build/wally ./dist/linux64/wally

docker kill --signal SIGKILL "$CONTAINER_NAME"
