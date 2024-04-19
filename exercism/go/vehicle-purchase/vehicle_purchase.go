package purchase

import (
	"fmt"
	"strings"
)

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if kind == "car" || kind == "truck" {
		return true
	}
	return false
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	result := strings.Compare(option1, option2)
	message := "%s is clearly the better choice."

	if result < 0 {
		return fmt.Sprintf(message, option1)
	}
	return fmt.Sprintf(message, option2)
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	if age < 3 {
		return originalPrice * .80
	} else if age >= 3 && age < 10 {
		return originalPrice * .70
	}
	return originalPrice * .50
}
