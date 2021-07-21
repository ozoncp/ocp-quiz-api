package utils

import "errors"

// Batches Разделение на слайса на батчи - исходный слайс конвертировать
// в слайс слайсов - чанки одинкового размера (кроме последнего)
func Batches(slice []int, size int) ([][]int, error) {
	if size <= 0 {
		return nil, errors.New("batch size must be less than zero")
	}
	result := make([][]int, 0)
	lenSlice := len(slice)
	chunkCount := lenSlice / size
	// Целые чанки - у котороых размер батча равен заданному
	for i := 0; i < chunkCount; i++ {
		start := i * size
		end := size + start
		result = append(result, slice[start:end])
	}
	// Добавить остаток, неполный чанк, если он есть
	n := lenSlice - lenSlice%size
	if n != lenSlice {
		result = append(result, slice[n:])
	}
	return result, nil
}
