package strain

// Ints abstracts a collection of numbers.
type Ints []int

// Lists abstracts a collection of lists of numbers.
type Lists [][]int

// Strings abstracts a collection of strings.
type Strings []string

// Keep returns filtered values that match a provided function.
func (in Ints) Keep(fn func(int) bool) (out Ints) {
	for _, i := range in {
		if fn(i) {
			out = append(out, i)
		}
	}
	return
}

// Discard returns filtered values that match a provided function.
func (in Ints) Discard(fn func(int) bool) (out Ints) {
	for _, i := range in {
		if !fn(i) {
			out = append(out, i)
		}
	}
	return
}

// Keep returns filtered values that match a provided function.
func (in Lists) Keep(fn func([]int) bool) (out Lists) {
	for _, i := range in {
		if fn(i) {
			out = append(out, i)
		}
	}
	return
}

// Keep returns filtered values that match a provided function.
func (in Strings) Keep(fn func(string) bool) (out Strings) {
	for _, i := range in {
		if fn(i) {
			out = append(out, i)
		}
	}
	return
}
