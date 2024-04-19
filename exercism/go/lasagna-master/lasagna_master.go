package lasagna

import "slices"

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) int {
	if time == 0 {
		time = 2
	}
	return len(layers) * time
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
	for _, v := range layers {
		if v == "noodles" {
			noodles += 50
		} else if v == "sauce" {
			sauce += 0.2
		}
	}
	return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) {
	for _, friendIngredient := range friendsList {
		for _, myIngredient := range myList {
			if !slices.Contains(friendsList, myIngredient) {
				myList[len(myList)-1] = friendIngredient
				break
			}
		}
	}
}

// TODO: define the 'ScaleRecipe()' functions
func ScaleRecipe(quantities []float64, portions int) []float64 {
	temp := make([]float64, len(quantities))
	copy(temp, quantities)
	for i := range temp {
		temp[i] = temp[i] * (float64(portions) / 2)
	}
	return temp
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
