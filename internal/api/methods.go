package api

import (
	"context"

	"github.com/ozoncp/ocp-quiz-api/internal/models"
	"github.com/ozoncp/ocp-quiz-api/internal/producer"
	"github.com/ozoncp/ocp-quiz-api/internal/utils"
	ocpQuizApi "github.com/ozoncp/ocp-quiz-api/pkg/ocp-quiz-api"

	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	tLog "github.com/opentracing/opentracing-go/log"
	"github.com/rs/zerolog/log"
)

func (s *api) CreateQuiz(ctx context.Context, req *ocpQuizApi.CreateQuizV1Request) (*ocpQuizApi.CreateQuizV1Response, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CreateQuiz")
	defer span.Finish()

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

	s.producer.Send(producer.NewEvent(ctx, quizId, producer.CreateEvent, err))
	s.metrics.IncCreate(1, "CreateQuiz")
	return &ocpQuizApi.CreateQuizV1Response{QuizId: quizId}, nil
}

func (s api) DescribeQuiz(ctx context.Context, req *ocpQuizApi.DescribeQuizV1Request) (*ocpQuizApi.DescribeQuizV1Response, error) {
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

	return &ocpQuizApi.DescribeQuizV1Response{Quiz: &ocpQuizApi.Quiz{
		Id:          found.Id,
		ClassroomId: found.ClassroomId,
		UserId:      found.UserId,
		Link:        found.Link,
	}}, nil
}

func (s *api) ListQuiz(ctx context.Context, req *ocpQuizApi.ListQuizV1Request) (*ocpQuizApi.ListQuizV1Response, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Info().
		Uint64("Limit", req.Limit).
		Uint64("Offset", req.Offset).
		Msg("ListQuizV1")

	quizes, _ := s.repo.ListEntities(ctx, req.Limit, req.Offset)

	resp := make([]*ocpQuizApi.Quiz, len(quizes))

	for i, r := range quizes {
		resp[i] = &ocpQuizApi.Quiz{
			Id:          r.Id,
			UserId:      r.UserId,
			ClassroomId: r.ClassroomId,
			Link:        r.Link,
		}
	}

	return &ocpQuizApi.ListQuizV1Response{
		Quizes:      resp,
		CurrentPage: req.Offset,
	}, nil
}

func (s *api) RemoveQuiz(ctx context.Context, req *ocpQuizApi.RemoveQuizV1Request) (*ocpQuizApi.RemoveQuizV1Response, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "RemoveQuiz")
	defer span.Finish()

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Info().
		Uint64("Id", req.QuizId).
		Msg("RemoveQuizV1")

	found := s.repo.RemoveEntity(ctx, req.QuizId)

	s.producer.Send(producer.NewEvent(ctx, req.QuizId, producer.DeleteEvent, nil))
	s.metrics.IncRemove(1, "RemoveQuiz")
	return &ocpQuizApi.RemoveQuizV1Response{Found: found}, nil
}

// MultiCreateQuiz  Creates new Quizes and returns this IDs
func (s *api) MultiCreateQuiz(ctx context.Context, req *ocpQuizApi.MultiCreateQuizV1Request) (*ocpQuizApi.MultiCreateQuizV1Response, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "RemoveQuiz")
	defer span.Finish()

	log.Info().Int("Count", len(req.Quizes)).Msg("MultiCreateRequestV1")

	quizes := make([]models.Quiz, len(req.Quizes))

	for i, req := range req.Quizes {
		quizes[i] = models.Quiz{
			Id:          0,
			ClassroomId: req.ClassroomId,
			UserId:      req.UserId,
			Link:        req.Link,
		}
	}
	batches, err := utils.SplitToBulks(quizes, s.batchSize)
	if err != nil {
		return nil, status.Error(codes.Aborted, err.Error())
	}

	ids := make([]uint64, 0, len(quizes))
	for _, batch := range batches {
		batchIds, err := s.writeRequestsBatch(ctx, batch)
		if err != nil {
			return nil, status.Error(codes.Canceled, err.Error())
		}
		ids = append(ids, batchIds...)
	}

	s.metrics.IncCreate(uint(len(quizes)), "MultiCreateQuiz")
	return &ocpQuizApi.MultiCreateQuizV1Response{
		QuizesIds: ids,
	}, nil
}

func (s *api) writeRequestsBatch(ctx context.Context, batch []models.Quiz) ([]uint64, error) {
	childSpan, childCtx := opentracing.StartSpanFromContext(ctx, "MultiCreateRequestV1Batch")
	childSpan.LogFields(tLog.Int("batch_size", len(batch)))

	defer childSpan.Finish()

	ids, err := s.repo.AddEntities(childCtx, batch)
	if err != nil {
		log.Error().
			Err(err).
			Msgf("Failed to save requests")

		return nil, err
	}
	for _, id := range ids {
		s.producer.Send(producer.NewEvent(ctx, id, producer.CreateEvent, nil))
	}
	return ids, nil
}

func (s *api) UpdateQuiz(ctx context.Context, req *ocpQuizApi.UpdateQuizV1Request) (*ocpQuizApi.UpdateQuizV1Response, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UpdateQuiz")
	defer span.Finish()

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Info().
		Uint64("Id", req.QuizId).
		Msg("UpdateQuiz")

	quiz := models.Quiz{
		Id:          req.QuizId,
		ClassroomId: req.ClassroomId,
		UserId:      req.UserId,
		Link:        req.Link,
	}

	updated, err := s.repo.UpdateEntity(ctx, quiz)
	if err != nil {
		return &ocpQuizApi.UpdateQuizV1Response{Updated: false}, err
	}

	s.producer.Send(producer.NewEvent(ctx, req.QuizId, producer.UpdateEvent, nil))
	s.metrics.IncUpdate(1, "UpdateQuiz")
	return &ocpQuizApi.UpdateQuizV1Response{Updated: updated}, nil
}
