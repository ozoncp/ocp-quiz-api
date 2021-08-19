package api

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/rs/zerolog/log"

	ocp_quiz_api "github.com/ozoncp/ocp-quiz-api/pkg/ocp-quiz-api"
)

func (s *api) CreateQuiz(ctx context.Context, req *ocp_quiz_api.CreateQuizV1Request) (*ocp_quiz_api.CreateQuizV1Response, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Info().
		Str("Link", req.Link).
		Uint64("ClassroomId", req.ClassroomId).
		Uint64("UserId", req.UserId).
		Msg("CreateQuizV1")

	return &ocp_quiz_api.CreateQuizV1Response{QuizId: 1}, nil
}

func (s api) DescribeQuiz(ctx context.Context, req *ocp_quiz_api.DescribeQuizV1Request) (*ocp_quiz_api.DescribeQuizV1Response, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Info().
		Uint64("id", req.QuizId).
		Msg("DescribeQuizV1")

	return &ocp_quiz_api.DescribeQuizV1Response{Quiz: &ocp_quiz_api.Quiz{
		Id:          1,
		ClassroomId: 2,
		UserId:      3,
		Link:        "123",
	}}, nil
}

func (s *api) ListQuiz(ctx context.Context, req *ocp_quiz_api.ListQuizV1Request) (*ocp_quiz_api.ListQuizV1Response, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Info().
		Uint64("Limit", req.Limit).
		Uint64("Offset", req.Offset).
		Msg("ListQuizV1")

	quizes := []*ocp_quiz_api.Quiz{{
		Id:          1,
		ClassroomId: 2,
		UserId:      3,
		Link:        "Link",
	}}

	return &ocp_quiz_api.ListQuizV1Response{
		Quizes:      quizes,
		CurrentPage: 1,
		IsLast:      true,
	}, nil
}

func (s *api) RemoveQuiz(ctc context.Context, req *ocp_quiz_api.RemoveQuizV1Request) (*ocp_quiz_api.RemoveQuizV1Response, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Info().
		Uint64("Id", req.QuizId).
		Msg("RemoveQuizV1")

	return &ocp_quiz_api.RemoveQuizV1Response{Found: true}, nil
}
