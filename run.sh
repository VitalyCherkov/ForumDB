#!/bin/sh

docker stop forumdb
docker rm forumdb
docker build -t=forumdb .
docker run -p 5000:5000 --name forumdb forumdb
