package gigasecond

import "time"

const (
	//Gigasecond is 10^9 (1,000,000,000) seconds
	Gigasecond = 1e9 * time.Second
)

// AddGigasecond given a moment, determine the moment that would be after a gigasecond has passed.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(Gigasecond)
}
