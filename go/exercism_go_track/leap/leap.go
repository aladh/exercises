// Package leap provides a function to calculate leap years
package leap

// IsLeapYear returns true if the year given is a leap year, otherwise it returns false
func IsLeapYear(year int) bool {
	return isDivisibleBy4(year) && (!isDivisibleBy100(year) || isDivisibleBy400(year))
}

func isDivisibleBy4(n int) bool {
	return n%4 == 0
}

func isDivisibleBy100(n int) bool {
	return n%100 == 0
}

func isDivisibleBy400(n int) bool {
	return n%400 == 0
}
