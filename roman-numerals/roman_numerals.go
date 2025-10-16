package romannumerals

import (
	"errors"
	"strings"
)

var nums = []int{1000, 500, 100, 50, 10, 5, 1}

/**

M	  	D	C	L	X	V	I
1000	500	100	50	10	5	1

4 （IV）、9 （IX）、40 （XL）、90 （XC）、400 （CD） 和 900 （CM）

*/

type RomanNumeral struct {
	Value  int
	Symbol string
}

var romanNumeral = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ToRomanNumeral(input int) (string, error) {
	if input <= 0 || input >= 4000 {
		return "", errors.New("")
	}

	var res strings.Builder
	for _, numeral := range romanNumeral {
		for input >= numeral.Value {
			res.WriteString(numeral.Symbol)
			input -= numeral.Value
		}
	}

	return res.String(), nil
}
