#!/bin/bash
# gluten-slave deploy

#create tmp folder
rm -rf ./tmp
mkdir -p ./tmp 

#prepare tmp folder 
cd ./tmp
cp -rf ../Dockerfile ./
cp -rf ../../migrations ./

#build gluten-slave
go build ../../../slave

#deploy gluten-slave
docker build -t gluten/slave .

docker rm -f gluten-slave
docker run --name gluten-slave -d gluten/slave -p 8888 -p 8889 /opt/gluten-slave/slave -m 8888 -r 8890
