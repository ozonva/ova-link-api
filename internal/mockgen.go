package internal

//go:generate ../bin/mockgen -destination=./mocks/repo_mock.go -package=mocks github.com/ozonva/ova-link-api/internal/repo Repo
//go:generate ../bin/mockgen -destination=./mocks/producer_mock.go -package=mocks github.com/ozonva/ova-link-api/internal/kafka Producer
//go:generate ../bin/mockgen -destination=./mocks/prom_mock.go -package=mocks github.com/ozonva/ova-link-api/internal/metrics Metrics
