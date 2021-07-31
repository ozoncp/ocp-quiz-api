package flusher

import (
	"github.com/ozoncp/ocp-quiz-api/internal/models"
	"github.com/ozoncp/ocp-quiz-api/internal/repo"
	"github.com/ozoncp/ocp-quiz-api/internal/utils"
)

// Flusher - интерфейс для сброса задач в хранилище
type Flusher interface {
	Flush(entities []models.Quiz) ([]models.Quiz, error)
}

type flusher struct {
	chunkSize  int
	entityRepo repo.Repo
}

// Flush - записать объекты в Repo по батчам
func (f *flusher) Flush(entities []models.Quiz) ([]models.Quiz, error) {
	batches, err := utils.SplitToBulks(entities, f.chunkSize)
	if err != nil {
		return nil, err
	}
	for _, chunk := range batches {
		if err := f.entityRepo.AddEntities(chunk); err != nil {
			return nil, err
		}
	}

	return make([]models.Quiz, 0), nil
}

// NewFlusher возвращает Flusher с поддержкой батчевого сохранения
func NewFlusher(
	chunkSize int,
	quizRepo repo.Repo,
) Flusher {
	return &flusher{
		chunkSize:  chunkSize,
		entityRepo: quizRepo,
	}
}
