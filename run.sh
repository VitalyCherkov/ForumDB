#!/bin/sh

docker stop forumdb
docker rm forumdb
docker build -t=forumdb .
docker run -p 5000:5000 -p 5432:5432 --rm --name forumdb forumdb
