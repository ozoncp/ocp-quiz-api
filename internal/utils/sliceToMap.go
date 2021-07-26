package utils

import "github.com/ozoncp/ocp-quiz-api/internal/models"

func SliceToMap(input []models.Quiz) map[uint64]models.Quiz {
	result := make(map[uint64]models.Quiz)
	for _, q := range input {
		result[q.Id] = q
	}
	return result
}
