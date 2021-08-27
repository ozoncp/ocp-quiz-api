package flusher

import (
	"context"

	"github.com/ozoncp/ocp-quiz-api/internal/mocks"
	"github.com/ozoncp/ocp-quiz-api/internal/models"
	"github.com/ozoncp/ocp-quiz-api/internal/repo"
	"github.com/ozoncp/ocp-quiz-api/internal/utils"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

var _ = Describe("Flush into repo by", func() {
	var (
		ctrl     *gomock.Controller
		mockRepo *mocks.MockRepo
		entities []models.Quiz
		ctx      context.Context
	)
	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockRepo(ctrl)
		entities = []models.Quiz{
			{1, 1, 1, "some_link1"},
			{21, 1, 2, "some_link21"},
			{31, 1, 3, "some_link31"},
			{41, 1, 4, "some_link41"},
		}
		ctx = context.Background()
	})

	AfterEach(func() { ctrl.Finish() })

	Context("batches for", func() {
		It("1 batches", func() {
			f := NewFlusher(4, mockRepo)
			mockRepo.EXPECT().AddEntities(gomock.Any(), gomock.Any()).Times(1)

			f.Flush(ctx, entities)

		})
		It("many batches", func() {
			f := NewFlusher(2, mockRepo)
			mockRepo.EXPECT().AddEntities(gomock.Any(), gomock.Any()).Times(2)

			f.Flush(ctx, entities)

		})
		It("many batches with tail", func() {
			f := NewFlusher(3, mockRepo)
			mockRepo.EXPECT().AddEntities(gomock.Any(), gomock.Any()).Times(2)

			f.Flush(ctx, entities)
		})
		It("1 batches where over max batch", func() {
			f := NewFlusher(5, mockRepo)
			mockRepo.EXPECT().AddEntities(gomock.Any(), gomock.Any()).Times(1)

			f.Flush(ctx, entities)
		})
		It("error from chunkSize", func() {
			f := NewFlusher(0, mockRepo)
			mockRepo.EXPECT().AddEntities(gomock.Any(), gomock.Any()).Times(0)

			_, err := f.Flush(ctx, entities)
			gomega.Expect(err).Should(gomega.Equal(utils.ErrBatchSizeZero))
		})
		It("error, returned from Repo", func() {
			f := NewFlusher(1, mockRepo)
			mockRepo.EXPECT().AddEntities(gomock.Any(), gomock.Any()).Return(nil, repo.ErrCannotAddEntity)

			res, err := f.Flush(ctx, entities)
			gomega.Expect(err).Should(gomega.Equal(repo.ErrCannotAddEntity))
			gomega.Expect(res).To(gomega.BeNil())
		})
	})
})
