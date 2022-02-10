package clock

import "fmt"

const (
	// Mh minutes per hour
	Mh int = 60
	// Hd hours per day
	Hd int = 24
	// Md minutes per day
	Md int = Mh * Hd
)

// Clock abstracts the concept of a clock.
type Clock struct {
	Hours, Minutes int
}

// New returns a clock.
func New(h, m int) Clock {
	n := (h*Mh + m) % Md
	if n < 0 {
		n += Md
	}

	return Clock{Hours: n / Mh, Minutes: n % Mh}
}

// String returns the text representation of a clock.
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.Hours, c.Minutes)
}

// Add returns the clock updated with added minutes.
func (c Clock) Add(m int) Clock {
	return New(c.Hours, c.Minutes+m)
}

// Subtract returns the clock updated with substracted minutes.
func (c Clock) Subtract(m int) Clock {
	return New(c.Hours, c.Minutes-m)
}
