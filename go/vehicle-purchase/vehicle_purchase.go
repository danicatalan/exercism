package purchase

const (
	VEHICLE_TYPE_1      string  = "car"
	VEHICLE_TYPE_2      string  = "truck"
	RECOMMENDATION_TEXT string  = " is clearly the better choice."
	AGE_THRESHOLD_1     float64 = 3
	AGE_THRESHOLD_2     float64 = 10
	COST_THRESHOLD_0    float64 = 0.8
	COST_THRESHOLD_1    float64 = 0.7
	COST_THRESHOLD_2    float64 = 0.5
)

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return kind == VEHICLE_TYPE_1 || kind == VEHICLE_TYPE_2
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in dictionary order.
func ChooseVehicle(option1, option2 string) string {
	if option1 < option2 {
		return option1 + RECOMMENDATION_TEXT
	}
	return option2 + RECOMMENDATION_TEXT
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	if age >= AGE_THRESHOLD_2 {
		return originalPrice * COST_THRESHOLD_2
	}
	if age >= AGE_THRESHOLD_1 {
		return originalPrice * COST_THRESHOLD_1
	}
	return originalPrice * COST_THRESHOLD_0
}
