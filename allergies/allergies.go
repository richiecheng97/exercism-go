package allergies

var allergenList = []string{
	"eggs",
	"peanuts",
	"shellfish",
	"strawberries",
	"tomatoes",
	"chocolate",
	"pollen",
	"cats",
}

func Allergies(allergies uint) []string {
	var result []string
	for i, allergen := range allergenList {
		if allergies&(1<<i) != 0 {
			result = append(result, allergen)
		}
	}
	return result
}

func AllergicTo(allergies uint, allergen string) bool {
	for i, a := range allergenList {
		if a == allergen {
			return allergies&(1<<i) != 0
		}
	}
	return false
}
