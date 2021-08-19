package main

import (
	"bytes"
	"context"
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"

	"github.com/ozoncp/ocp-quiz-api/internal/api"
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

func run() error {
	listener, err := net.Listen("tcp", grpcPort)
	if err != nil {
		return errors.Wrapf(err, "failed to start tcp listener on %s", grpcPort)
	}

	server := grpc.NewServer()
	ocp_quiz_api.RegisterOcpQuizApiServiceServer(server, api.NewOcpQuizApiService())

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

func main() {
	go runJSON()

	if err := run(); err != nil {
		log.Fatal().
			Msgf(err.Error())
	}
}
