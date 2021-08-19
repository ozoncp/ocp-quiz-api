package api

import ocp_quiz_api "github.com/ozoncp/ocp-quiz-api/pkg/ocp-quiz-api"

type api struct {
	ocp_quiz_api.UnimplementedOcpQuizApiServiceServer
}

func NewOcpQuizApiService() ocp_quiz_api.OcpQuizApiServiceServer {
	return &api{}
}
