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
