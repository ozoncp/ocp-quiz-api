package utils

// https://stackoverflow.com/questions/10485743/contains-method-for-a-slice
func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// Filter Фильтрация по захордкоженному списку
// фильтрует входной слайс по критерию отсутствия элемента в захардкоженном списке
func Filter(input []int) []int {
	hardcodedSlice := []int{1, 2, 3}
	result := make([]int, 0)
	for _, i := range input {
		if ok := contains(hardcodedSlice, i); !ok {
			result = append(result, i)
		}
	}
	return result
}
