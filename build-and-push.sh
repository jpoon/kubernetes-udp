#!/bin/bash

set -e

. var.env

echo server container is: ${SERVER_CONTAINER}
echo client container is: ${CLIENT_CONTAINER}

docker build -f ./server/Dockerfile -t ${SERVER_CONTAINER} .

docker build -f ./client/Dockerfile -t ${CLIENT_CONTAINER} .

docker push ${SERVER_CONTAINER}

docker push ${CLIENT_CONTAINER}