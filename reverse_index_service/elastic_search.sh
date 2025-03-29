#!/bin/sh
docker run -p 8080:9200 -it -d -e "discovery.type=single-node" elasticsearch:7.9.3
