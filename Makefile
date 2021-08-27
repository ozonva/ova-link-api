run:
	go run ./cmd/ova-link-api/main.go

run-config:
	go run ./cmd/ova-link-api/main.go config

run-server:
	go run ./cmd/ova-link-api/main.go server

build:
	go build ./cmd/ova-link-api/

test-internal:
	./bin/ginkgo ./internal/...

generate-mocks:
	go generate ./internal/mockgen.go

BIN_PATH = $(PWD)/bin
install:
	mkdir -p bin && \
 	go get -d github.com/onsi/ginkgo && \
 	go get -d google.golang.org/protobuf/cmd/protoc-gen-go && \
    go get -d google.golang.org/grpc/cmd/protoc-gen-go-grpc && \
    go get -d google.golang.org/grpc && \
    go get -d github.com/golang/protobuf/proto && \
    GOBIN=$(BIN_PATH) go install google.golang.org/protobuf/cmd/protoc-gen-go && \
    GOBIN=$(BIN_PATH) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc && \
    GOBIN=$(BIN_PATH) go install google.golang.org/grpc && \
    GOBIN=$(BIN_PATH) go install github.com/golang/protobuf/proto && \
    GOBIN=$(BIN_PATH) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc && \
 	GOBIN=$(BIN_PATH) go install github.com/onsi/ginkgo/ginkgo@v1.16.4 && \
 	GOBIN=$(BIN_PATH) go install github.com/golang/mock/mockgen@v1.6.0

proto:
	PATH=$(PATH):$(BIN_PATH) protoc -I vendor.protogen \
        --go_out=pkg/ --go_opt=paths=import \
        --go-grpc_out=pkg/ --go-grpc_opt=paths=import \
        --proto_path=./api link.proto