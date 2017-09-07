#!/bin/bash
# gluten-master deploy

docker network create gluten 
docker rm -f gluten-master
docker run --name gluten-db --network gluten -e POSTGRES_PASSWORD=1 -e POSTGRES_DB=gluten -p 5432:5432 -d postgres
docker run --name gluten-master --network gluten  -p 8888:8888 -p 8889:8889 -d gluten/master

