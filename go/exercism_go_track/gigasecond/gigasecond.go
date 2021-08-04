// Package gigasecond provides a function to add a gigasecond to a time
package gigasecond

import "time"

const gigaSecond = time.Duration(1e9 * time.Second)

// AddGigasecond adds one gigasecond to the time
func AddGigasecond(t time.Time) time.Time {
	return t.Add(gigaSecond)
}
