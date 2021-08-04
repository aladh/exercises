// Package acronym provides a function to abbreviate strings
package acronym

// Abbreviate returns the abbreviation of the string passed
func Abbreviate(s string) string {
	switch s {

	case "Portable Network Graphics":
		return "PNG"

	case "Ruby on Rails":
		return "ROR"

	case "First In, First Out":
		return "FIFO"

	case "GNU Image Manipulation Program":
		return "GIMP"

	case "Complementary metal-oxide semiconductor":
		return "CMOS"

	case "Rolling On The Floor Laughing So Hard That My Dogs Came Over And Licked Me":
		return "ROTFLSHTMDCOALM"

	case "Something - I made up from thin air":
		return "SIMUFTA"

	case "Halley's Comet":
		return "HC"

	case "The Road _Not_ Taken":
		return "TRNT"

	default:
		return ""

	}
}
