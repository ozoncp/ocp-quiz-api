package api

import (
	"github.com/ozoncp/ocp-quiz-api/internal/metrics"
	"github.com/ozoncp/ocp-quiz-api/internal/producer"
	"github.com/ozoncp/ocp-quiz-api/internal/repo"
	ocp_quiz_api "github.com/ozoncp/ocp-quiz-api/pkg/ocp-quiz-api"
)

type api struct {
	ocp_quiz_api.UnimplementedOcpQuizApiServiceServer
	repo      repo.Repo
	batchSize int
	m         metrics.MetricsReporter
	p         producer.Producer
}

func NewOcpQuizApiService(r repo.Repo, batchSize int, m metrics.MetricsReporter, p producer.Producer) *api {
	return &api{
		repo:      r,
		batchSize: batchSize,
		m:         m,
		p:         p,
	}
}
