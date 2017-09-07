#!/bin/bash
# gluten-master deploy

#create tmp folder
rm -rf ./tmp
mkdir -p ./tmp 

#prepare tmp folder 
cd ./tmp
cp -rf ../Dockerfile ./
cp -rf ../../migrations ./
cp -rf ../../../master-config.yml ./

#build gluten-master
go build ../../../master/backend

#deploy gluten-master
docker build -t gluten/master .
