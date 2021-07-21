package utils

// Swap Обратный ключ - происходит конвертация отображения
// (“ключ-значение“) в отображение (“значение-ключ“)
func Swap(input map[string]string) map[string]string {
	result := make(map[string]string)
	for key, value := range input {
		result[value] = key
	}
	return result
}
