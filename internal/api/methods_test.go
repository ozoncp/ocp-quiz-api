package api

import (
	"errors"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"golang.org/x/net/context"

	"github.com/ozoncp/ocp-quiz-api/internal/mocks"
	"github.com/ozoncp/ocp-quiz-api/internal/models"
	"github.com/ozoncp/ocp-quiz-api/internal/repo"
	ocp_quiz_api "github.com/ozoncp/ocp-quiz-api/pkg/ocp-quiz-api"
)

type AnyContext struct {
}

func (c AnyContext) Matches(val interface{}) bool {
	if _, ok := val.(context.Context); ok {
		return true
	}
	return false
}

func (c AnyContext) String() string {
	return "this is not the droid you're looking for"
}

var _ = Describe("Methods", func() {

	var (
		reqApi       *api
		mockRepo     *mocks.MockRepo
		mockCtrl     *gomock.Controller
		ctx          context.Context
		mockMetrics  *mocks.MockMetricsReporter
		mockProducer *mocks.MockProducer
		anyCtx       AnyContext
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockRepo(mockCtrl)
		mockMetrics = mocks.NewMockMetricsReporter(mockCtrl)
		mockProducer = mocks.NewMockProducer(mockCtrl)
		ctx = context.Background()
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})
	Context("Request", func() {
		JustBeforeEach(func() {
			reqApi = NewOcpQuizApiService(mockRepo, 5, mockMetrics, mockProducer)
			ctx = context.Background()
			anyCtx = AnyContext{}
		})
		It("Create without errors", func() {
			newQuizId := uint64(19)

			mockProducer.EXPECT().
				Send(gomock.Any())

			mockMetrics.EXPECT().
				IncCreate(uint(1), "CreateQuiz")

			mockRepo.EXPECT().
				AddEntity(anyCtx, gomock.Any()).
				Return(newQuizId, nil).
				MaxTimes(1).
				MinTimes(1)
			resp, err := reqApi.CreateQuiz(
				ctx, &ocp_quiz_api.CreateQuizV1Request{UserId: 10, ClassroomId: 11, Link: "link"},
			)

			Expect(resp).
				To(Equal(&ocp_quiz_api.CreateQuizV1Response{
					QuizId: newQuizId,
				}))

			Expect(err).ToNot(HaveOccurred())
		})
		It("Create with errors", func() {
			mockRepo.EXPECT().
				AddEntity(anyCtx, gomock.Any()).
				Return(uint64(0), errors.New("some error")).
				MaxTimes(1).
				MinTimes(1)
			resp, err := reqApi.CreateQuiz(
				ctx, &ocp_quiz_api.CreateQuizV1Request{UserId: 10, ClassroomId: 11, Link: "link"},
			)

			Expect(resp).To(BeNil())

			Expect(err.Error()).To(Equal("rpc error: code = Aborted desc = some error"))
		})
		Context("MultiCreate", func() {
			request := []*ocp_quiz_api.CreateQuizV1Request{
				{ClassroomId: 100, UserId: 1000, Link: "one"},
				{ClassroomId: 200, UserId: 2000, Link: "two"},
				{ClassroomId: 300, UserId: 3000, Link: "three"},
			}
			It("without errors", func() {
				expIds := []uint64{1, 2, 3}

				mockProducer.EXPECT().
					Send(gomock.Any()).
					MinTimes(1)

				mockMetrics.EXPECT().
					IncCreate(uint(len(expIds)), "MultiCreateQuiz")

				mockRepo.EXPECT().
					AddEntities(anyCtx, gomock.Any()).
					Return(expIds, nil).
					MaxTimes(1).
					MinTimes(1)
				resp, err := reqApi.MultiCreateQuiz(
					ctx, &ocp_quiz_api.MultiCreateQuizV1Request{Quizes: request},
				)

				expected := ocp_quiz_api.MultiCreateQuizV1Response{QuizesIds: expIds}
				Expect(err).ToNot(HaveOccurred())
				Expect(resp).To(Equal(&expected))

			})
			It("with errors", func() {
				expError := errors.New("test")
				mockRepo.EXPECT().
					AddEntities(anyCtx, gomock.Any()).
					Return(nil, expError).
					MaxTimes(1).
					MinTimes(1)
				resp, err := reqApi.MultiCreateQuiz(
					ctx, &ocp_quiz_api.MultiCreateQuizV1Request{Quizes: request},
				)

				Expect(err.Error()).To(Equal("rpc error: code = Canceled desc = test"))
				Expect(resp).To(BeNil())

			})
		})

		It("List without error", func() {
			offset, limit := uint64(1), uint64(3)
			requests := []models.Quiz{
				{1, 100, 1000, "one"},
				{2, 200, 2000, "two"},
				{3, 300, 3000, "three"},
			}

			mockRepo.EXPECT().
				ListEntities(ctx, limit, offset).
				Return(requests, nil).
				MaxTimes(1).
				MinTimes(1)
			resp, err := reqApi.ListQuiz(
				ctx, &ocp_quiz_api.ListQuizV1Request{
					Offset: offset, Limit: limit,
				},
			)

			req := make([]*ocp_quiz_api.Quiz, 0, len(requests))
			for _, r := range requests {
				req = append(req, &ocp_quiz_api.Quiz{
					Id:          r.Id,
					UserId:      r.UserId,
					ClassroomId: r.ClassroomId,
					Link:        r.Link,
				})
			}

			Expect(resp).
				To(Equal(&ocp_quiz_api.ListQuizV1Response{
					Quizes:      req,
					CurrentPage: offset,
				}))

			Expect(err).ToNot(HaveOccurred())
		})
		It("Describe existing request", func() {
			quiz := models.Quiz{
				Id:          1,
				UserId:      10,
				ClassroomId: 1000,
				Link:        "one",
			}
			mockRepo.EXPECT().
				DescribeEntity(ctx, quiz.Id).
				Return(&quiz, nil).
				MaxTimes(1).
				MinTimes(1)
			resp, err := reqApi.DescribeQuiz(
				ctx, &ocp_quiz_api.DescribeQuizV1Request{
					QuizId: quiz.Id,
				},
			)

			Expect(resp).
				To(Equal(&ocp_quiz_api.DescribeQuizV1Response{
					Quiz: &ocp_quiz_api.Quiz{
						Id:          quiz.Id,
						UserId:      quiz.UserId,
						ClassroomId: quiz.ClassroomId,
						Link:        quiz.Link,
					},
				}))

			Expect(err).ToNot(HaveOccurred())
		})

		It("Describe not existing request", func() {
			mockRepo.EXPECT().
				DescribeEntity(ctx, uint64(1)).
				Return(nil, repo.NotExists).
				MaxTimes(1).
				MinTimes(1)
			resp, err := reqApi.DescribeQuiz(
				ctx, &ocp_quiz_api.DescribeQuizV1Request{
					QuizId: uint64(1),
				},
			)

			Expect(resp).To(BeNil())
			Expect(err.Error()).To(Equal("rpc error: code = NotFound desc = entity does not exist"))
		})

		Context("Update test", func() {
			quiz := models.Quiz{
				Id:          1,
				ClassroomId: 1,
				UserId:      1,
				Link:        "one",
			}

			update := func(expectUpdate bool) {
				mockProducer.EXPECT().
					Send(gomock.Any()).
					MinTimes(1)

				mockMetrics.EXPECT().
					IncUpdate(uint(1), "UpdateQuiz")

				mockRepo.EXPECT().
					UpdateEntity(anyCtx, quiz).
					Return(expectUpdate, nil).
					MaxTimes(1).
					MinTimes(1)

				resp, err := reqApi.UpdateQuiz(
					ctx, &ocp_quiz_api.UpdateQuizV1Request{
						QuizId:      quiz.Id,
						UserId:      quiz.UserId,
						ClassroomId: quiz.ClassroomId,
						Link:        quiz.Link,
					},
				)

				Expect(resp).
					To(Equal(&ocp_quiz_api.UpdateQuizV1Response{
						Updated: expectUpdate,
					}))

				Expect(err).ToNot(HaveOccurred())
			}
			It("exist", func() {
				update(true)
			})
			It("not exist", func() {
				update(false)
			})
			It("invalid arguments", func() {
				resp, err := reqApi.UpdateQuiz(
					ctx, &ocp_quiz_api.UpdateQuizV1Request{
						QuizId:      0,
						UserId:      0,
						ClassroomId: 0,
						Link:        "quiz.Link",
					},
				)

				Expect(resp).To(BeNil())
				Expect(err.Error()).To(Equal("rpc error: code = InvalidArgument desc = invalid UpdateQuizV1Request.QuizId: value must be greater than 0"))
			})
			It("err by repo", func() {
				testErr := errors.New("test")
				mockRepo.EXPECT().
					UpdateEntity(anyCtx, quiz).
					Return(false, testErr).
					MaxTimes(1).
					MinTimes(1)

				resp, err := reqApi.UpdateQuiz(
					ctx, &ocp_quiz_api.UpdateQuizV1Request{
						QuizId:      quiz.Id,
						UserId:      quiz.UserId,
						ClassroomId: quiz.ClassroomId,
						Link:        quiz.Link,
					},
				)

				Expect(resp).To(Equal(&ocp_quiz_api.UpdateQuizV1Response{
					Updated: false,
				}))
				Expect(err.Error()).To(Equal("test"))
			})
		})

		Context("Remove test", func() {
			remove := func(expectFound bool) {
				requestId := uint64(1)

				mockProducer.EXPECT().
					Send(gomock.Any()).
					MinTimes(1)

				mockMetrics.EXPECT().
					IncRemove(uint(1), "RemoveQuiz")

				mockRepo.EXPECT().
					RemoveEntity(anyCtx, requestId).
					Return(expectFound).
					MaxTimes(1).
					MinTimes(1)

				resp, err := reqApi.RemoveQuiz(
					ctx, &ocp_quiz_api.RemoveQuizV1Request{
						QuizId: requestId,
					},
				)

				Expect(resp).
					To(Equal(&ocp_quiz_api.RemoveQuizV1Response{
						Found: expectFound,
					}))

				Expect(err).ToNot(HaveOccurred())
			}
			It("Remove existing entity", func() {
				remove(true)
			})
			It("Remove non-existing entity", func() {
				remove(false)
			})

		})

		Context("Invalid param", func() {
			It("ClassroomId in CreateQuizV1Request", func() {
				_, err := reqApi.CreateQuiz(
					ctx, &ocp_quiz_api.CreateQuizV1Request{},
				)
				Expect(err.Error()).To(Equal("rpc error: code = InvalidArgument desc = invalid CreateQuizV1Request.ClassroomId: value must be greater than 0"))

			})
			It("UserId in CreateQuizV1Request", func() {
				_, err := reqApi.CreateQuiz(
					ctx, &ocp_quiz_api.CreateQuizV1Request{UserId: 0, ClassroomId: 1},
				)
				Expect(err.Error()).To(Equal("rpc error: code = InvalidArgument desc = invalid CreateQuizV1Request.UserId: value must be greater than 0"))

			})
			It("QuizId in RemoveQuizV1Request", func() {
				_, err := reqApi.RemoveQuiz(
					ctx, &ocp_quiz_api.RemoveQuizV1Request{QuizId: 0},
				)
				Expect(err.Error()).To(Equal("rpc error: code = InvalidArgument desc = invalid RemoveQuizV1Request.QuizId: value must be greater than 0"))

			})

			It("QuizId in DescribeQuizV1Request", func() {
				_, err := reqApi.DescribeQuiz(
					ctx, &ocp_quiz_api.DescribeQuizV1Request{},
				)
				Expect(err.Error()).To(Equal("rpc error: code = InvalidArgument desc = invalid DescribeQuizV1Request.QuizId: value must be greater than 0"))

			})

			It("Limit in ListQuizV1Request", func() {
				_, err := reqApi.ListQuiz(
					ctx, &ocp_quiz_api.ListQuizV1Request{Limit: 0},
				)
				Expect(err.Error()).To(Equal("rpc error: code = InvalidArgument desc = invalid ListQuizV1Request.Limit: value must be greater than or equal to 1"))

			})
		})

	})
})
