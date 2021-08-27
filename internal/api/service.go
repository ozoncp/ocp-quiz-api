package api

import (
	"github.com/ozoncp/ocp-quiz-api/internal/repo"
	ocp_quiz_api "github.com/ozoncp/ocp-quiz-api/pkg/ocp-quiz-api"
)

type api struct {
	ocp_quiz_api.UnimplementedOcpQuizApiServiceServer
	repo repo.Repo
}

func NewOcpQuizApiService(r repo.Repo) *api {
	return &api{
		repo: r,
	}
}
