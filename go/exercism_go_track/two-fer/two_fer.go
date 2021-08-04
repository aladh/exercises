// Package twofer implements a function that returns a string
// based on the name given
package twofer

import "fmt"

// ShareWith returns "One for X, one for me." where X is the
// name given, or "you" if no name is given
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
