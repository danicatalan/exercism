package space

var (
	orbitalPeriod = map[Planet]float64{
		"Earth":   1.0,
		"Mercury": 0.2408467,
		"Venus":   0.61519726,
		"Mars":    1.8808158,
		"Jupiter": 11.862615,
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 164.79132,
	}
)

const (
	orbitalPeriodInSeconds float64 = 31557600
)

// Planet is an object that orbits stars
type Planet string

// Age calculates how old is somebody in relative years for a specific planet
func Age(ageInSeconds float64, planet Planet) (ageInYears float64) {
	ageInYears = ageInSeconds / (orbitalPeriod[planet] * orbitalPeriodInSeconds)
	return
}
