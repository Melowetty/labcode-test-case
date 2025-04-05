package utils

func ConvertArray[I, T any](initial []I, mapFunction func(I) T) []T {
	result := make([]T, 0, len(initial))
	for _, value := range initial {
		convertedValue := mapFunction(value)
		result = append(result, convertedValue)
	}

	return result
}
