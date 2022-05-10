package booking

import (
	"time"
	"fmt"
)

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	time, _ := time.Parse("1/2/2006 15:04:05", date)
	return time
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	currentTime := time.Now()

	formattedTimeParam,_ := time.Parse("January 2, 2006 15:04:05", date)

	return formattedTimeParam.Before(currentTime)
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	formattedTimeParam,_ := time.Parse("Monday, January 2, 2006 15:04:05", date)

	return formattedTimeParam.Hour() >= 12 && formattedTimeParam.Hour() <= 17
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	time, _ := time.Parse("1/2/2006 15:04:05", date)

	return fmt.Sprintf("You have an appointment on %v, %v %v, %v, at %v:%v.", time.Weekday(), time.Month(), time.Day(), time.Year(), time.Hour(), time.Minute())
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
