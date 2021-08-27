package api

import (
	"github.com/ozoncp/ocp-quiz-api/internal/metrics"
	"github.com/ozoncp/ocp-quiz-api/internal/producer"
	"github.com/ozoncp/ocp-quiz-api/internal/repo"
	ocpQuizApi "github.com/ozoncp/ocp-quiz-api/pkg/ocp-quiz-api"
)

type api struct {
	ocpQuizApi.UnimplementedOcpQuizApiServiceServer
	repo      repo.Repo
	batchSize int
	metrics   metrics.MetricsReporter
	producer  producer.Producer
}

func NewOcpQuizApiService(r repo.Repo, batchSize int, m metrics.MetricsReporter, p producer.Producer) *api {
	return &api{
		repo:      r,
		batchSize: batchSize,
		metrics:   m,
		producer:  p,
	}
}
