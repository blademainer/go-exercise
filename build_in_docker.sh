#!/usr/bin/env bash
file="$1"
docker run --rm -v `pwd`:/tmp -it golang bash -c "go get -v github.com/Masterminds/glide ; go get gopkg.in/yaml.v2; cd /go/src/github.com/Masterminds/glide ; git checkout e73500c735917e39a8b782e0632418ab70250341 ; go install ; cd /tmp; glide install --update-vendored; env GOOS=linux GOARCH=arm go build -o /tmp/bin/$file /tmp/$file"
#docker run --rm -v `pwd`:/tmp -it golang env GOOS=linux GOARCH=arm go build -o /tmp/bin/hello /tmp/hello.go
