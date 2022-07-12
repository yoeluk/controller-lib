#!/usr/bin/env bash

go build -o dbcontroller -ldflags "-X google.golang.org/protobuf/reflect/protoregistry.conflictPolicy=warn"
docker build -t dbcontroller .