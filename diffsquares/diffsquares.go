package diffsquares

// SquareOfSum returns (1 + .. + n)^2 where n number of natural numbers
func SquareOfSum(n int) int {
	sum := (1 + n) * n / 2
	return sum * sum
}

// SumOfSquares returns 1^2 + .. + n^2 where n number of natural numbers
func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

// Returns difference between SquareOfSum and SumOfSquares
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
