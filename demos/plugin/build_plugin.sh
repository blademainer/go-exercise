#!/usr/bin/env bash
go build  -buildmode=plugin -o hello.so ./provider/

mv hello.so ../../