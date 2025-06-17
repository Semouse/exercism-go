// Package leap contains function to define if year is leap year.
package leap

// IsLeapYear returns true if year is leap year, otherwise returns false.
func IsLeapYear(year int) bool {
	if year%100 == 0 {
		return year%400 == 0
	}

	return year%4 == 0
}
