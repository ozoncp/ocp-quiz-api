package main

import (
	"bytes"
	"context"
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/Shopify/sarama"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"

	"github.com/ozoncp/ocp-quiz-api/internal/api"
	"github.com/ozoncp/ocp-quiz-api/internal/db"
	"github.com/ozoncp/ocp-quiz-api/internal/metrics"
	"github.com/ozoncp/ocp-quiz-api/internal/producer"
	repo "github.com/ozoncp/ocp-quiz-api/internal/repo"
	ocp_quiz_api "github.com/ozoncp/ocp-quiz-api/pkg/ocp-quiz-api"
)

func FileReader(files []string) {
	fileReader := func(fname string) (string, error) {
		file, err := os.Open(fname)
		if err != nil {
			return "", err
		}

		defer file.Close()

		data := new(bytes.Buffer)
		if _, err = data.ReadFrom(file); err != nil {
			return "", err
		}

		return data.String(), nil
	}

	for _, file := range files {
		_, err := fileReader(file)
		if err != nil {
			fmt.Printf("fileReader error: %v", err)
		}
	}
}

const (
	grpcPort           = ":8081"
	grpcServerEndpoint = ":8081"
	httpPort           = ":8083"
)

func getProducer() producer.Producer {
	brokers := []string{"localhost:9094"}

	cfg := sarama.NewConfig()
	cfg.Producer.Partitioner = sarama.NewRandomPartitioner
	cfg.Producer.RequiredAcks = sarama.WaitForAll
	cfg.Producer.Return.Successes = true
	p, err := sarama.NewSyncProducer(brokers, cfg)
	kafkaTopic := "ocp-quiz-api"

	if err != nil {
		log.Panic().Msgf("failed to connect to Kafka brokers: %v", err)
	}

	return producer.NewProducer(kafkaTopic, p)
}

func run(r repo.Repo, m metrics.MetricsReporter, p producer.Producer) error {
	listener, err := net.Listen("tcp", grpcPort)
	if err != nil {
		return errors.Wrapf(err, "failed to start tcp listener on %s", grpcPort)
	}

	server := grpc.NewServer()
	ocp_quiz_api.RegisterOcpQuizApiServiceServer(server, api.NewOcpQuizApiService(r, 5, m, p))

	log.Info().
		Str("address", listener.Addr().String()).
		Msg("serving gRPC...")

	if err := server.Serve(listener); err != nil {
		return errors.Wrap(err, "failed to serve gRPC")
	}

	return nil
}

func runJSON() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	if err := ocp_quiz_api.RegisterOcpQuizApiServiceHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts); err != nil {
		log.Panic().
			Msgf("failed to register API handler from endpoint '%s': %v", grpcServerEndpoint, err)
	}

	log.Info().
		Str("address", httpPort).
		Msg("serving REST...")

	if err := http.ListenAndServe(httpPort, mux); err != nil {
		log.Panic().
			Msgf("failed to serve JSON endpoint: %v", err)
	}
}

func runMetrics() {
	http.Handle("/metrics", promhttp.Handler())
	if err := http.ListenAndServe(":9100", nil); err != nil {
		log.Panic().Msgf("metrics don't started: %v", err)
	}
	log.Info().Int("port", 9100).Msg("metrics started")
}

func main() {
	dsn := os.Getenv("DB_DSN")
	database := db.Connect(dsn)
	defer database.Close()

	r := repo.NewRepo(database)
	m := metrics.NewMetricsReporter()
	p := getProducer()

	go runJSON()
	go runMetrics()

	if err := run(r, m, p); err != nil {
		log.Fatal().
			Msgf(err.Error())
	}
}
