package cars

const (
	MINUTES_IN_HOUR = 60
	GROUP_SIZE      = 10
	INDIVIDUAL_COST = 10000
	GROUP_COST      = 95000
)

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (successRate / 100)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	carsPerHour := CalculateWorkingCarsPerHour(productionRate, successRate)
	return int(carsPerHour / MINUTES_IN_HOUR)
}

// CalculateCost works out the cost of producing the given number of cars
func CalculateCost(carsCount int) uint {
	groups, rest := carsCount/GROUP_SIZE, carsCount%GROUP_SIZE
	return uint(groups*GROUP_COST + rest*INDIVIDUAL_COST)
}
