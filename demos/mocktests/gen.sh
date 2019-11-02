#!/usr/bin/env bash

 mockgen -source repo.go  -package mock -destination repo_mock.go
