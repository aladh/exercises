// Package bob provides a function to emulate a lackadaisical teenager
package bob

import (
	"regexp"
	"strings"
)

// Hey returns a response based on the string given
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

	switch {
	case isQuestion(remark) && isYelling(remark):
		return "Calm down, I know what I'm doing!"
	case isQuestion(remark):
		return "Sure."
	case isYelling(remark):
		return "Whoa, chill out!"
	case isBlank(remark):
		return "Fine. Be that way!"
	default:
		return "Whatever."
	}
}

func isQuestion(remark string) bool {
	return strings.HasSuffix(remark, "?")
}

func isYelling(remark string) bool {
	match, _ := regexp.MatchString("[A-Z]", remark)

	return match && strings.ToUpper(remark) == remark
}

func isBlank(remark string) bool {
	return len(remark) == 0
}
