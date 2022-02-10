package summultiples

//SumMultiples given a number, find the sum of all the unique multiples of particular numbers
//up to but not including that number.
func SumMultiples(limit int, numbers ...int) (sum int) {
	multiples := make(map[int]int)

	for _, n := range numbers {
		if n != 0 {
			for i := n; i < limit; i += n {
				multiples[i]++
			}
		}
	}

	for m := range multiples {
		sum += m
	}

	return sum
}
