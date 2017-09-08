#!/bin/bash
# gluten-master local deploy

# current dir 
WORK_DIR=$(pwd)

# network 
docker network create gluten

# DB
docker run --name gluten-db --network gluten -e POSTGRES_PASSWORD=1 -e POSTGRES_DB=gluten -p 5432:5432 -d postgres

# nginx & UI
cd ./tmp
cp -rf ../nginx.conf ./
cd ..
cp -rf ../../master/frontend/dist ./
docker run --name gluten-nginx --network gluten -d -p 80:80 -v $WORK_DIR/nginx.conf:/etc/nginx/nginx.conf -v $WORK_DIR/dist/:/usr/share/nginx/html  nginx || docker restart gluten-nginx

# gluten-master
docker rm -f gluten-master
docker run --name gluten-master --network gluten  -p 8888:8888 -p 8889:8889 -d gluten/master

