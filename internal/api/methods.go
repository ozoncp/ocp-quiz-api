package api

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/rs/zerolog/log"

	"github.com/ozoncp/ocp-quiz-api/internal/models"
	"github.com/ozoncp/ocp-quiz-api/internal/utils"
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

// MultiCreateQuiz  Creates new Quizes and returns this IDs
func (s *api) MultiCreateQuiz(ctx context.Context, req *ocp_quiz_api.MultiCreateQuizV1Request) (*ocp_quiz_api.MultiCreateQuizV1Response, error) {
	log.Info().Int("Count", len(req.Quizes)).Msg("MultiCreateRequestV1")

	quizes := make([]models.Quiz, 0, len(req.Quizes))

	for _, req := range req.Quizes {
		quizes = append(quizes, models.Quiz{
			Id:          0,
			ClassroomId: req.ClassroomId,
			UserId:      req.UserId,
			Link:        req.Link,
		})
	}
	batches, err := utils.SplitToBulks(quizes, s.batchSize)
	if err != nil {
		return nil, status.Error(codes.Aborted, err.Error())
	}

	ids := make([]uint64, 0, len(quizes))

	for _, batch := range batches {
		batchIds, err := s.repo.AddEntities(ctx, batch)

		if err != nil {
			return nil, status.Error(codes.Canceled, err.Error())
		}

		for _, i := range batchIds {
			ids = append(ids, i)
		}
	}

	return &ocp_quiz_api.MultiCreateQuizV1Response{
		QuizesIds: ids,
	}, nil
}
