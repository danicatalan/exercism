/*
Package diffsquares A naive implementation using a loop to iterate from 1 to n
has a time complexity O(n), while using a mathematical formula reduces the time complexity to O(1).
*/
package diffsquares

//SquareOfSum takes an number and returns the square of the sum of first n numbers.
func SquareOfSum(n int) int {
	sum := n * (n + 1) / 2
	return sum * sum
}

//SumOfSquares takes a number and returns the sum of squares of first n numbers. .
func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

//Difference takes a number and returns the difference between its square of sum and its sum of squares.
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
