package resistorcolortrio

import "fmt"

/*
*
black: 0  黑色：0
brown: 1  棕色：1
red: 2  红色：2
orange: 3  橙色：3
yellow: 4  黄色：4
green: 5  绿色：5
blue: 6  蓝色：6
violet: 7  紫罗兰色：7
grey: 8  灰色：8
white: 9  白色：9
*/
var mp = map[string]int{
	"black":  0,
	"brown":  1,
	"red":    2,
	"orange": 3,
	"yellow": 4,
	"green":  5,
	"blue":   6,
	"violet": 7,
	"grey":   8,
	"white":  9,
}

var units = []string{"ohms", "kiloohms", "megaohms", "gigaohms"}

// Label describes the resistance value given the colors of a resistor.
// The label is a string with a resistance value with an unit appended
// (e.g. "33 ohms", "470 kiloohms").
func Label(colors []string) string {
	if len(colors) > 3 {
		colors = colors[:3]
	}

	value := mp[colors[0]]*10 + mp[colors[1]]
	multiplier := mp[colors[2]]
	for i := 0; i < multiplier; i++ {
		value *= 10
	}

	return formatValue(value)
}

func formatValue(value int) string {
	index := 0
	for value >= 1000 && index < len(units)-1 {
		value /= 1000
		index++
	}

	return fmt.Sprintf("%d %s", value, units[index])
}
