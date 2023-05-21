#!/bin/bash

echo "Building image for redis server..."
cd ./redis-server 
docker build -t redis-server .

echo "Building image for Go client..."
cd ../client-go
docker build -t client-go .

cd ..
