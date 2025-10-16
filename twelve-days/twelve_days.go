package twelve

import "fmt"

var strs = []string{"twelve Drummers Drumming,", "eleven Pipers Piping,", "ten Lords-a-Leaping,", "nine Ladies Dancing,", "eight Maids-a-Milking,", "seven Swans-a-Swimming,", "six Geese-a-Laying,", "five Gold Rings,", "four Calling Birds,", "three French Hens,", "two Turtle Doves,", "a Partridge in a Pear Tree."}
var nums = []string{"first", "second", "third", "fourth", "fifth", "sixth", "seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth"}

/**
On the twelfth day of Christmas my true love gave to me: twelve Drummers Drumming, eleven Pipers Piping, ten Lords-a-Leaping, nine Ladies Dancing, eight Maids-a-Milking, seven Swans-a-Swimming, six Geese-a-Laying, five Gold Rings, four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree.
*/

func Verse(n int) string {
	var res = fmt.Sprintf("On the %s day of Christmas my true love gave to me:", nums[n-1])
	for i := n; i > 0; i-- {
		str := " " + strs[12-i]
		if 12-i == 11 && n != 1 {
			str = " and " + strs[12-i]
		}
		res += str
	}
	return res
}

func Song() string {
	var res string
	for i := 1; i <= 12; i++ {
		res += Verse(i)
		if i != 12 {
			res += "\n"
		}
	}
	return res
}
