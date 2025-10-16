package resistorcolorduo

import "strings"

// Value should return the resistance value of a resistor with a given colors.
func Value(colors []string) int {
	return ColorCode(colors[0])*10 + ColorCode(colors[1])
}

// Colors returns the list of all colors.
func Colors() []string {
	return []string{
		"black",
		"brown",
		"red",
		"orange",
		"yellow",
		"green",
		"blue",
		"violet",
		"grey",
		"white",
	}
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
	color = strings.ToLower(color)
	colors := Colors()
	for i := 0; i < len(colors); i++ {
		if colors[i] == color {
			return i
		}
	}
	return -1
}
