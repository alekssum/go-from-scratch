#!/bin/sh

docker build -t go-mysql . && \
docker run -d -p 3306:3306 --name=go-mysql go-mysql
