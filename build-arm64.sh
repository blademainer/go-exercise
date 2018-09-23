#!/bin/bash

cur_dir="`pwd`"
mkdir work-dir
cp web-crawler.go web-crawler.yaml web-crawler-docker work-dir/
#docker run --rm -v `pwd`/work-dir/:/work-dir:rw -it golang go get gopkg.in/yaml.v2 ; env GOOS=linux GOARCH=arm64 go build -o /work-dir/web-crawler /work-dir/web-crawler.go
#docker run --rm -v `pwd`/work-dir/:/work-dir:rw -it golang "go get gopkg.in/yaml.v2"
cd work-dir
env GOOS=linux GOARCH=arm go build -o web-crawler web-crawler.go
docker build -f web-crawler-docker -t web-crawler:v1.0.0 .
cd "$cur_dir"
rm -fr work-dir
