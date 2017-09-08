#!/bin/bash
# gluten-slave build

#create tmp folder
rm -rf ./tmp
mkdir -p ./tmp 

#prepare tmp folder 
cd ./tmp
cp -rf ../Dockerfile ./
cp -rf ../../migrations ./
cp -rf ../../../slave-config.yml ./

#build gluten-slave
go build ../../../slave

#build gluten-slave image
docker build -t gluten/slave .
