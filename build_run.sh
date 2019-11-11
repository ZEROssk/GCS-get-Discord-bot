#!/bin/bash
docker build ./ -t ggcsdb/zero:latest
docker run -it --restart=always --name ggcsdb_zero ggcsdb/zero:latest
