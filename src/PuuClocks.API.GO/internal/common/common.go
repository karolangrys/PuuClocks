package common

func Chunk[T any](m []T, howManyChunks int) [][]T {
	var res [][]T

	n := len(m) / howManyChunks

	for i := 0; i < howManyChunks; i++ {
		minIndex := i * n
		maxIndex := (i + 1) * n

		res = append(res, m[minIndex:maxIndex])
	}

	return res
}

func Contains[T comparable](item T, arr []T) bool {
	for _, i := range arr {
		if i == item {
			return true
		}
	}
	return false
}
