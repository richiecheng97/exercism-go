package meetup

import "time"

// Define the WeekSchedule type here.

type WeekSchedule string

const (
	First  WeekSchedule = "first"
	Second WeekSchedule = "second"
	Third  WeekSchedule = "third"
	Fourth WeekSchedule = "fourth"
	Last   WeekSchedule = "last"
	Teenth WeekSchedule = "teenth"
)

func Day(wSched WeekSchedule, wDay time.Weekday, month time.Month, year int) int {
	switch wSched {
	case Teenth:
		for day := 13; day <= 19; day++ {
			tt := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
			if tt.Weekday() == wDay {
				return day
			}
		}
	case Last:
		last := time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC).Day()
		for day := last; day >= 1; day-- {
			tt := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
			if tt.Weekday() == wDay {
				return day
			}
		}
	default:
		occurrence := 0
		switch wSched {
		case First:
			occurrence = 1
		case Second:
			occurrence = 2
		case Third:
			occurrence = 3
		case Fourth:
			occurrence = 4
		}

		count := 0
		for day := 1; day <= 31; day++ {
			tt := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
			if tt.Month() != month {
				break
			}
			if tt.Weekday() == wDay {
				count++
				if count == occurrence {
					return day
				}
			}
		}
	}
	return 0
}
