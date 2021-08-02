package repo

import (
	"errors"

	"github.com/ozoncp/ocp-quiz-api/internal/models"
)

var ErrCannotAddEntity = errors.New("cannot add entity")

// Repo - интерфейс хранилища для сущности Quiz
type Repo interface {
	AddEntities(entities []models.Quiz) error
	ListEntities(limit, offset uint64) ([]models.Quiz, error)
	DescribeEntity(entityId uint64) (*models.Quiz, error)
}
