
GOPATH:=$(shell go env GOPATH)


.PHONY: proto
proto:
	protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/helloworld.proto
	mv github.com/nicocesar/helloworldwithmemory/helloworld.pb.* proto

.PHONY: build
build: proto

	go build -o helloworld-srv *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t helloworld-srv:latest
