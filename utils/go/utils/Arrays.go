package utils

func SplitToChunks[T comparable](arr []T, chunkSize int) [][]T {
	var chunks [][]T
	for i := 0; i < len(arr); i += chunkSize {
		end := i + 3
		if end > len(arr) {
			end = len(arr)
		}
		chunks = append(chunks, arr[i:end])
	}
	return chunks
}

func SumOfArray(arr []int) (sum int) {
	for _, number := range arr {
		sum += number
	}
	return sum
}

func DefaultValuedSlice[T comparable](size int, val T) []T {
	retArr := make([]T, size)
	for i := 0; i < size; i++ {
		retArr[i] = val
	}
	return retArr
}
