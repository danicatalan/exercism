package reverse

//Reverse takes a string and returns the reversed one.
func Reverse(input string) (output string) {
	for _, l := range input {
		output = string(l) + output
	}
	return
}
