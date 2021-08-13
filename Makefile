run:
	go run ./cmd/ova-link-api/main.go

run-config:
	go run ./cmd/ova-link-api/main.go config

build:
	go build ./cmd/ova-link-api/

test-internal:
	go test ./internal/...