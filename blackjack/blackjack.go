package blackjack

var mp = map[string]int{
	"ace":   11,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"ten":   10,
	"jack":  10,
	"queen": 10,
	"king":  10,
	"other": 0,
}

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	return mp[card]
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	num := ParseCard(card1) + ParseCard(card2)
	dealer := ParseCard(dealerCard)
	switch {
	case num == 22:
		return "P"
	case num == 21 && dealer < 10:
		return "W"
	case num == 21 && dealer >= 10:
		return "S"
	case num >= 17 && num <= 20:
		return "S"
	case num >= 12 && num <= 26 && dealer < 7:
		return "S"
	case num >= 12 && num <= 26 && dealer >= 7:
		return "H"
	default:
		return "H"
	}
}
