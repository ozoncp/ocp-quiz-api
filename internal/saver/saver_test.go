package saver

import (
	"time"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"

	"github.com/ozoncp/ocp-quiz-api/internal/mocks"
	"github.com/ozoncp/ocp-quiz-api/internal/models"
)

var _ = Describe("Saver", func() {
	var (
		mockCtrl    *gomock.Controller
		mockFlusher *mocks.MockFlusher
		s           Saver
		timeout     time.Duration
		entities    []models.Quiz
		capacity    uint
	)
	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockFlusher = mocks.NewMockFlusher(mockCtrl)
		timeout = 100 * time.Millisecond

		entities = []models.Quiz{
			{1, 1, 41, "link_one"},
			{2, 2, 42, "link_two"},
			{3, 3, 43, "link_three"},
			{4, 4, 44, "link_four"},
			{5, 5, 45, "link_five"},
			{6, 6, 46, "link_six"},
		}
	})
	AfterEach(func() {
		mockCtrl.Finish()
	})
	It("flushes all entities", func() {
		capacity = uint(len(entities) / 2)
		s = NewSaver(capacity, mockFlusher, timeout)
		mockFlusher.EXPECT().Flush(gomock.Any()).MinTimes(1)
		defer s.Close()
		for _, entity := range entities {
			s.Save(entity)
		}
	})
	It("flushes all entities after close", func() {
		capacity = 0
		s = NewSaver(capacity, mockFlusher, timeout)
		mockFlusher.EXPECT().Flush(entities).MinTimes(1)
		defer s.Close()
		for _, entity := range entities {
			s.Save(entity)
		}
	})
	It("flushes all entities", func() {
		capacity = uint(len(entities))
		s = NewSaver(capacity, mockFlusher, timeout)
		mockFlusher.EXPECT().Flush(gomock.Any()).MinTimes(1)
		for _, entity := range entities {
			s.Save(entity)
		}
		s.Close()

		s.Save(entities[0])
		s.Save(entities[0])
		s.Save(entities[0])
	})
	It("successfully stopped", func() {
		capacity = uint(len(entities))
		s = NewSaver(capacity, mockFlusher, timeout)
		mockFlusher.EXPECT().Flush([]models.Quiz{}).Times(1)
		s.Close()
		for _, entity := range entities {
			s.Save(entity)
		}
	})
	It("flushes entities by timeout", func() {
		capacity = uint(len(entities))
		timeout = time.Millisecond
		s = NewSaver(capacity, mockFlusher, timeout)
		mockFlusher.EXPECT().Flush(gomock.Any()).MinTimes(1)

		for _, entity := range entities {
			s.Save(entity)
		}
		time.Sleep(50 * time.Millisecond)
	})
})
