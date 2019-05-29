#!/bin/bash
docker build ./ -t ggcsdb/zero:latest
docker run --name ggcsdb_zero --rm ggcsdb/zero:latest
