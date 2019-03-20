#!/usr/bin/env sh
echo "building the app"
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o phil .
echo "app built, creating Docker image"
docker build -t phil -f Dockerfile.scratch .
echo "docker image created"