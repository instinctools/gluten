#!/bin/bash
# gluten-slave local deploy

docker rm -f gluten-slave
docker run --name gluten-slave --network gluten -p 7000:7000 -d gluten/slave
