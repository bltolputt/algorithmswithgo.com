package module01

// Sum will sum up all of the numbers passed
// in and return the result
func Sum(numbers []int) int {
	if len(numbers) == 0 || numbers == nil {
		return 0
	}
	total := 0
	for _, v := range numbers {
		total += v
	}
	return total
}
