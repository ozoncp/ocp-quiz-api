package api

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/rs/zerolog/log"

	"github.com/ozoncp/ocp-quiz-api/internal/models"
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

	quiz := models.Quiz{
		Id:          0,
		ClassroomId: req.ClassroomId,
		UserId:      req.UserId,
		Link:        req.Link,
	}
	quizId, err := s.repo.AddEntity(ctx, quiz)
	if err != nil {
		return nil, status.Error(codes.Aborted, err.Error())
	}

	return &ocp_quiz_api.CreateQuizV1Response{QuizId: quizId}, nil
}

func (s api) DescribeQuiz(ctx context.Context, req *ocp_quiz_api.DescribeQuizV1Request) (*ocp_quiz_api.DescribeQuizV1Response, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Info().
		Uint64("id", req.QuizId).
		Msg("DescribeQuizV1")

	found, err := s.repo.DescribeEntity(ctx, req.QuizId)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &ocp_quiz_api.DescribeQuizV1Response{Quiz: &ocp_quiz_api.Quiz{
		Id:          found.Id,
		ClassroomId: found.ClassroomId,
		UserId:      found.UserId,
		Link:        found.Link,
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

	quizes, _ := s.repo.ListEntities(ctx, req.Limit, req.Offset)

	resp := make([]*ocp_quiz_api.Quiz, 0, len(quizes))

	for _, r := range quizes {
		resp = append(resp, &ocp_quiz_api.Quiz{
			Id:          r.Id,
			UserId:      r.UserId,
			ClassroomId: r.ClassroomId,
			Link:        r.Link,
		})
	}

	return &ocp_quiz_api.ListQuizV1Response{
		Quizes:      resp,
		CurrentPage: req.Offset,
	}, nil
}

func (s *api) RemoveQuiz(ctx context.Context, req *ocp_quiz_api.RemoveQuizV1Request) (*ocp_quiz_api.RemoveQuizV1Response, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Info().
		Uint64("Id", req.QuizId).
		Msg("RemoveQuizV1")

	found := s.repo.RemoveEntity(ctx, req.QuizId)

	return &ocp_quiz_api.RemoveQuizV1Response{Found: found}, nil
}
