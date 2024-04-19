// Package raindrops has methods to convert integers
package raindrops

import "strconv"

// Convert converts integers to PlingPlangPlong notation
func Convert(input int) string {
	result := ""
	if input%3 != 0 && input%5 != 0 && input%7 != 0 {
		return strconv.Itoa(input)
	} else {
		if input%3 == 0 {
			result += "Pling"
		}
		if input%5 == 0 {
			result += "Plang"
		}
		if input%7 == 0 {
			result += "Plong"
		}
	}
	return result
}
