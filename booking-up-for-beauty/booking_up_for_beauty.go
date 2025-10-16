package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	parsedTime, _ := time.ParseInLocation("1/2/2006 15:04:05", date, time.UTC)
	return parsedTime
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	parsedTime, _ := time.ParseInLocation("January 2, 2006 15:04:05", date, time.UTC)
	return parsedTime.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	parsedTime, _ := time.ParseInLocation("Monday, January 2, 2006 15:04:05", date, time.UTC)
	return parsedTime.Hour() >= 12 && parsedTime.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	parsedTime, _ := time.ParseInLocation("1/2/2006 15:04:05", date, time.UTC)
	return parsedTime.Format("You have an appointment on Monday, January 2, 2006, at 15:04.")
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	parsedtime, _ := time.ParseInLocation("2006-01-02 00:00:00 +0000 UTC", fmt.Sprintf("%d-09-15 00:00:00 +0000 UTC", time.Now().Year()), time.UTC)
	return parsedtime
}
