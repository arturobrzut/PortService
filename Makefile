IMAGE ?= quay.io/arturobrzut/portservice:1.0

all: grpc, test, build, docker-build

grpc:
	protoc -I ./proto  --go_out=./pkg/grpc --go-grpc_opt=require_unimplemented_servers=false  --go-grpc_out=./pkg/grpc   ./proto/*.proto

test:
	go test -v ./pkg/...

test-e2e:
	 go test -v ./e2e/...

build: fmt
	go build -o ./bin/port-service cmd/portservice/main.go

docker-build: test
	docker build -t ${IMAGE} .

docker-push: docker-build
	docker push ${IMAGE}

run: build
	./bin/port-service

fmt:
	go fmt ./...

vet:
	go vet ./...

install_proto:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1

.PHONY: all, grpc, test, test-e2e, build, docker-build, docker-push, run, fmt, vet, install_proto