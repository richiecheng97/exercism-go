package gross

var mp = map[string]int{
	"quarter_of_a_dozen": 3,
	"half_of_a_dozen":    6,
	"dozen":              12,
	"small_gross":        120,
	"gross":              144,
	"great_gross":        1728,
}

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return mp
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int, 0)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	unitValue, exists := units[unit]
	if !exists {
		return false
	}

	bill[item] += unitValue
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	if _, exists := bill[item]; !exists {
		return false
	}

	unitValue, unitExists := units[unit]
	if !unitExists {
		return false
	}

	newQuantity := bill[item] - unitValue
	if newQuantity < 0 {
		return false
	} else if newQuantity == 0 {
		delete(bill, item)
		return true
	} else {
		bill[item] = newQuantity
		return true
	}
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	value, ok := bill[item]
	if !ok {
		return 0, false
	}
	return value, true
}
