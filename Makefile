run:
	go run ./cmd/ova-link-api/main.go

run-config:
	go run ./cmd/ova-link-api/main.go config

build:
	go build ./cmd/ova-link-api/

test-internal:
	./bin/ginkgo ./internal/...

generate-mocks:
	go generate ./internal/mockgen.go

BIN_PATH = $(PWD)/bin
install-bin:
	mkdir -p bin && \
 	go get -d github.com/onsi/ginkgo && \
 	GOBIN=$(BIN_PATH) go install github.com/onsi/ginkgo/ginkgo && \
 	GOBIN=$(BIN_PATH) go install github.com/golang/mock/mockgen@v1.6.0