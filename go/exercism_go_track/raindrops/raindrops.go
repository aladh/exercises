package raindrops

import "strconv"

// Convert the given number to a string, based on the number's factors
func Convert(n int) string {
	var s string

	if isDivisibleBy(n, 3) {
		s += "Pling"
	}

	if isDivisibleBy(n, 5) {
		s += "Plang"
	}

	if isDivisibleBy(n, 7) {
		s += "Plong"
	}

	if len(s) == 0 {
		s += strconv.Itoa(n)
	}

	return s
}

func isDivisibleBy(n, factor int) bool {
	return n%factor == 0
}
