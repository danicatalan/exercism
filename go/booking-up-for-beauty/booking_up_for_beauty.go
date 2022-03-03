package booking

import (
	"time"
)

const (
	LAYOUT_SHORT       string     = "1/2/2006 15:04:05"
	LAYOUT_LONG        string     = "January 2, 2006 15:04:05"
	LAYOUT_FULL        string     = "Monday, January 2, 2006 15:04:05"
	LAYOUT_DESCRIPTION string     = "You have an appointment on Monday, January 2, 2006, at 15:04."
	AFTERNOON_STARTS   int        = 12
	AFTERNOON_ENDS     int        = 18
	ANNIVERSARY_MONTH  time.Month = 9
	ANNIVERSARY_DAY    int        = 15
)

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	for _, layout := range []string{
		LAYOUT_SHORT,
		LAYOUT_LONG,
		LAYOUT_FULL,
	} {
		t, err := time.Parse(layout, date)
		if err == nil {
			return t
		}
	}
	return time.Time{}
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	return Schedule(date).Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	return Schedule(date).Hour() >= AFTERNOON_STARTS && Schedule(date).Hour() < AFTERNOON_ENDS
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	return Schedule(date).Format(LAYOUT_DESCRIPTION)
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), ANNIVERSARY_MONTH, ANNIVERSARY_DAY, 0, 0, 0, 0, time.UTC)
}
