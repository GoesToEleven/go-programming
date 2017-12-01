package mymath

// Fib returns the nth number in the Fibonacci sequence.
// Fibonacci sequence adds up the previous two numbers:
// 0 1 1 2 3 5 8 13 21
func Fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
