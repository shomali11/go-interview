package countrotations

import "github.com/shomali11/go-interview/slices/rotations/minrotations"

// Count counts the number of rotations
func Count(array []int) int {
	minimumIndex := minrotations.Min(array)
	if minimumIndex == 0 {
		return 0
	}
	return len(array) - minimumIndex
}
