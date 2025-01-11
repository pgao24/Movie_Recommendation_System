#!/bin/sh
#This program build an image and start a container
TAG="${1:-latest}"
docker build -f dockerfile -t movie_recommendation_image:"$TAG" . || { echo "docker build failed" ; exit 1; } 
docker run -p 5080:8080 -v /Users/pgao/desktop/Movie_Project/go_code:/movie_recommendation -it --rm movie_recommendation_image:"$TAG" /bin/bash

