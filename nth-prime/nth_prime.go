package prime

//Nth takes a positive number and return the nth prime number in that position.
//Returns !ok if the given position is zero o negative.
func Nth(n int) (int, bool) {
	if n <= 0 {
		return 0, false
	}

	c := 0
	for p := 2; ; p++ {
		if isPrime(p) {
			c++
			if c == n {
				return p, true
			}
		}
	}
}

func isPrime(n int) bool {
	if n%2 == 0 {
		return n == 2
	}

	for i := 3; i*i < n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}
