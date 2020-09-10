# Helloworld Service

This is the Helloworld service with some silly changes to test the micro framework

## Overview

A simple helloworld service, with some memory of all the poeple that has been greeted.

## Usage

```
# Install micro v3
curl -fsSL https://install.m3o.com/micro | /bin/bash
# Install protoc (folow https://grpc.io/docs/protoc-installation/ )
# ...
# Install proto-gen-micro v3 (this is important! you might need to override v2 if already installed)
go get -u github.com/micro/micro/v3/cmd/protoc-gen-micro@master
export PATH=$PATH:~/go/bin

# run the server
micro server &

# run the service
make build
./helloworld-srv &

## call the service
micro call helloworld Helloworld.Call '{"name": "Alice"}'
micro call helloworld Helloworld.Call '{"name": "Bob"}'
micro call helloworld Helloworld.Call '{"name": "Charlie"}'
```
