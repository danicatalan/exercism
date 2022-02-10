package sieve

//Sieve finds all prime numbers up to a given limit.
func Sieve(limit int) []int {
	var list []bool
	for i := 0; i <= limit; i++ {
		list = append(list, true)
	}

	primes := []int{}
	for i := 0; i < len(list); i++ {
		if i < 2 {
			list[i] = false
		}
		if list[i] {
			primes = append(primes, i)
			for j := i + i; j < len(list); j += i {
				list[j] = false
			}
		}
	}

	return primes
}
