package resistorcolor

import "slices"

// Colors returns the list of all colors.
func Colors() []string {
	colorList := []string{"black", "brown", "red", "orange", "yellow", "green", "blue", "violet", "grey", "white"}
	return colorList
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
	colors := Colors()
	return slices.Index(colors, color)
}
