#!/bin/bash
# gluten-cli deploy

#create tmp folder
rm -rf ./tmp
mkdir -p ./tmp 

#prepare tmp folder 
cd ./tmp
cp -rf ../Dockerfile ./
cp -rf ../json ./

#build gluten-slave
go build ../../../cli

#deploy gluten-slave
docker build -t gluten/cli .

docker rm -f gluten-cli
docker run --name gluten-cli -d gluten/cli /opt/gluten-cli/cli -m 172.17.0.1:8888 -p /opt/gluten-cli/json