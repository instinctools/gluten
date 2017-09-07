#!/bin/bash
# gluten-master deploy

#create tmp folder
rm -rf ./tmp
mkdir -p ./tmp 

#prepare tmp folder 
cd ./tmp
cp -rf ../Dockerfile ./
cp -rf ../../migrations ./

#build gluten-master
go build ../../../master/backend

#deploy gluten-master
docker build -t gluten/master .

docker rm -f gluten-master
docker run --name gluten-master -p 8888:8888 -p 8889:8889 -d gluten/master /opt/gluten-master/master -r 8888 -w 8889
