package arraysandslices

func Sum(numbers [5]int) int {
	var result int
	for i := range len(numbers) {
		result += numbers[i]
	}
	return result
}
