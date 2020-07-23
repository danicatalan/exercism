package encode

import (
	"strconv"
	"strings"
	"unicode"
)

//RunLengthEncode takes a string and returns its run-length encoded string.
func RunLengthEncode(input string) string {
	var output strings.Builder
	var count int

	for i := range input {
		if i < len(input)-1 && input[i] == input[i+1] {
			count++
			continue
		}

		if count > 0 {
			output.WriteString(strconv.Itoa(count + 1))
		}

		output.WriteByte(input[i])
		count = 0
	}

	return output.String()
}

//RunLengthDecode takes a run-length encoded string and returns its original string.
func RunLengthDecode(input string) string {
	var output strings.Builder
	var count string

	for _, c := range input {
		if unicode.IsNumber(c) {
			count += string(c)
			continue
		}

		if count == "" {
			count = "1"
		}

		r, _ := strconv.Atoi(count)
		output.WriteString(strings.Repeat(string(c), r))
		count = ""
	}

	return output.String()
}
