// Package gigasecond calculates exact date and time after interval of thousand million seconds.
package gigasecond

import "time"

// AddGigasecond apply interval of thousand million second to passed date.
func AddGigasecond(t time.Time) time.Time {
	var gigasecond int64 = 1e+18
	duration := time.Duration(gigasecond)

	return t.Add(duration)
}
