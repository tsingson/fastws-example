#!/usr/bin/env bash

protoc  -I=/Users/qinshen/go/src -I=/usr/local/include  -I=./ --gofast_out=plugins=grpc:.  ./*.proto

protoc --lint_out=. *.proto

flatc --go --grpc ./*.fbs

flatc -g --gen-object-api --grpc ./*.fbs



flatc  --go --gen-object-api --gen-all  --gen-mutable --grpc  --gen-compare  --raw-binary ./*.fbs


flatc -s --gen-mutable ./*.fbs
