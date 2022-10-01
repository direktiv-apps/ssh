#!/bin/sh

docker build -t ssh . && docker run -p 9191:8080 ssh