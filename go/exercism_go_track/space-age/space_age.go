// Package space implements a function to calculate age on different planets
package space

// Planet : The name of a planet
type Planet string

const earthSeconds = 31557600

var orbitalRatios = map[Planet]float64{
	"Earth":   1,
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

// Age calculates age in years given age in seconds and the name of a planet
func Age(ageInSeconds float64, planet Planet) float64 {
	earthYears := ageInSeconds / earthSeconds

	return earthYears / orbitalRatios[planet]
}
