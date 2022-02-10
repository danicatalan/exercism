package accumulate

// Accumulate takes a collection and an operation to perform on each element of the collection,
// then returns a new collection containing the result of applying that operation to each element of the input collection.
func Accumulate(input []string, operation func(string) string) (output []string) {
	for _, i := range input {
		output = append(output, operation(i))
	}
	return
}
