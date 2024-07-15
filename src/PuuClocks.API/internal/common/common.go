package common

func Chunk[T any](m []T, howManyChunks int) [][]T {
	var res [][]T

	n := len(m) / howManyChunks

	for i := 0; i < howManyChunks; i++ {

		min := i * n
		max := (i + 1) * n

		res = append(res, m[min:max])

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
