package purchase
import "strings"

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return (strings.ToLower(kind) == "car" ||  strings.ToLower(kind) == "truck")
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	if (option1 > option2) {
		return option2 + " is clearly the better choice."
	}  else {
		return option1 + " is clearly the better choice."
	}
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var factor float64;

	if age >= 10 {
		factor = .5
	} else if age < 10 && age >= 3 {
		factor = .7
	} else {
		factor = .8
	}

	return float64(originalPrice) * factor
}
