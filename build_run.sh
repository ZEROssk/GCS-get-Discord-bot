#!/bin/bash
docker build ./ -t ggcsdb/zero:latest
docker run -it --name ggcsdb_zero --rm ggcsdb/zero:latest
