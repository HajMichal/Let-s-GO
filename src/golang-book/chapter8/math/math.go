package math

// The good practice of naming functions in go is every function that is visible outside should be wrote with first Capital letter.
func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}