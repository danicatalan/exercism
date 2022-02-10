package perfect

import "errors"

// Classification abstract the categories to classify numbers.
type Classification int

const (
	// ClassificationPerfect aliquot sum = number.
	ClassificationPerfect Classification = iota

	// ClassificationAbundant aliquot sum > number.
	ClassificationAbundant

	// ClassificationDeficient aliquot sum < number.
	ClassificationDeficient
)

var (
	// ErrOnlyPositive returned if the provided number is not positive.
	ErrOnlyPositive = errors.New("Number should be positive")
)

// Classify determines if a number is perfect, abundant, or deficient based on
// Nicomachus' classification scheme for natural numbers.
func Classify(n int64) (Classification, error) {
	if n <= 0 {
		return -1, ErrOnlyPositive
	}

	var sum int64
	for i := int64(1); i <= n/2; i++ {
		if n%i == 0 {
			sum += i
		}
	}

	if sum > n {
		return ClassificationAbundant, nil
	}
	if sum < n {
		return ClassificationDeficient, nil
	}
	return ClassificationPerfect, nil
}
