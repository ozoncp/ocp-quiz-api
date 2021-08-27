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

var _ = Describe("Methods", func() {
	var (
		reqApi   *api
		mockRepo *mocks.MockRepo
		mockCtrl *gomock.Controller
		ctx      context.Context
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockRepo(mockCtrl)
		ctx = context.Background()
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})
	Context("Request", func() {
		JustBeforeEach(func() {
			reqApi = NewOcpQuizApiService(mockRepo)
			ctx = context.Background()
		})
		It("Create without errors", func() {
			newQuizId := uint64(19)
			mockRepo.EXPECT().
				AddEntity(ctx, gomock.Any()).
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
				AddEntity(ctx, gomock.Any()).
				Return(uint64(0), errors.New("some error")).
				MaxTimes(1).
				MinTimes(1)
			resp, err := reqApi.CreateQuiz(
				ctx, &ocp_quiz_api.CreateQuizV1Request{UserId: 10, ClassroomId: 11, Link: "link"},
			)

			Expect(resp).To(BeNil())

			Expect(err.Error()).To(Equal("rpc error: code = Aborted desc = some error"))
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
		Context("Remove test", func() {
			remove := func(expectFound bool) {
				requestId := uint64(1)
				mockRepo.EXPECT().
					RemoveEntity(ctx, requestId).
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
