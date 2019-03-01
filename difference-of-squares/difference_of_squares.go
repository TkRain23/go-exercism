package diffsquares

// SquareOfSum computes (sum 1..i)**2
func SquareOfSum(i int) int {
	sum := i * (i + 1) / 2
	return sum * sum
}

// SumOfSquares computes sum (1..i)**2
func SumOfSquares(i int) int {
	return i * (i + 1) * (2*i + 1) / 6
}

// Difference computes SquareOfSum(i) - SumOfSquares(i)
func Difference(i int) int {
	return SquareOfSum(i) - SumOfSquares(i)
}
