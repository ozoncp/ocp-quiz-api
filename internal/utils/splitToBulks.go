package utils

import (
	"github.com/ozoncp/ocp-quiz-api/internal/models"
)

func SplitToBulks(input []models.Quiz, size int) ([][]models.Quiz, error) {
	if size <= 0 {
		return nil, ErrBatchSizeZero
	}
	result := make([][]models.Quiz, 0)
	lenSlice := len(input)
	chunkCount := lenSlice / size
	// Целые чанки - у котороых размер батча равен заданному
	for i := 0; i < chunkCount; i++ {
		start := i * size
		end := size + start
		result = append(result, input[start:end])
	}
	// Добавить остаток, неполный чанк, если он есть
	n := lenSlice - lenSlice%size
	if n != lenSlice {
		result = append(result, input[n:])
	}
	return result, nil
}
