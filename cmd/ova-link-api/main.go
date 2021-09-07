package main

import (
	"io"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/opentracing/opentracing-go"

	"github.com/uber/jaeger-client-go"

	"github.com/ozonva/ova-link-api/internal/metrics"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/ozonva/ova-link-api/internal/kafka"

	"github.com/jmoiron/sqlx"
	"github.com/ozonva/ova-link-api/internal/api"
	"github.com/ozonva/ova-link-api/internal/repo"
	"github.com/rs/zerolog"

	linkAPI "github.com/ozonva/ova-link-api/pkg/ova-link-api"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	jaegermetrics "github.com/uber/jaeger-lib/metrics"

	"google.golang.org/grpc"
)

func main() {
	const grpcPort = ":82"
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	db, err := sqlx.Open("pgx", "postgres://user_links:123456@localhost:5432/user_links?sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	producer, err := kafka.NewProducer([]string{"127.0.0.1:9093"})
	if err != nil {
		log.Fatalln(err)
	}

	go func() {
		http.Handle("/metrics", promhttp.Handler())
		err = http.ListenAndServe(":9200", nil)
		if err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	cfg := jaegercfg.Configuration{
		ServiceName: "ova-link-api",
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans: true,
		},
	}

	jLogger := jaegerlog.StdLogger
	jMetricsFactory := jaegermetrics.NullFactory

	tracer, closer, err := cfg.NewTracer(
		jaegercfg.Logger(jLogger),
		jaegercfg.Metrics(jMetricsFactory),
	)
	if err != nil {
		log.Fatalf("Can not create tracer, %s", err)
	}

	opentracing.SetGlobalTracer(tracer)
	defer func(io.Closer) {
		err := closer.Close()
		if err != nil {
			log.Fatalf("Can not close tracer, %s", err)
		}
	}(closer)

	handler := api.NewLinkAPI(
		repo.NewLinkRepo(db),
		zerolog.New(os.Stdout),
		producer,
		metrics.NewMetrics(),
	)
	linkAPI.RegisterLinkAPIServer(s, handler)

	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	return
}
