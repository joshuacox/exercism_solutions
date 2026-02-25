package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	gross_units := map[string]int{}
	gross_units["quarter_of_a_dozen"] = 3
	gross_units["half_of_a_dozen"] = 6
	gross_units["dozen"] = 12
	gross_units["small_gross"] = 120
	gross_units["gross"] = 144
	gross_units["great_gross"] = 1728
	return gross_units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	bill := map[string]int{}
	return bill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	// Return false if the given unit is not in the units map.
	// https://go.dev/blog/maps
	unitExists := units[unit]
	if unitExists == 0 {
		return false
	} else {
		itemExistsInBill := bill[item]
		if itemExistsInBill == 0 {
			// Otherwise add the item to the customer bill, indexed by 
			// the item name, then return true.
			bill[item] = units[unit]
			return true
		} else {
			// If the item is already present in the bill, increase its 
			// quantity by the amount that belongs to the provided unit.
			bill[item] = bill[item] + units[unit]
			return true
		}
	}
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	itemExistsInBill := bill[item]
	unitExists := units[unit]
	// Return false if the given item is not in the bill
	// Return false if the given unit is not in the units map.
	if itemExistsInBill == 0 || unitExists == 0 {
		return false
	}
	newVal := bill[item] - units[unit]
	// Return false if the new quantity would be less than 0.
	if newVal < 0 {
		return false
	} 
	// If the new quantity is 0, completely remove the item from 
	// the bill then return true.
	if newVal == 0 {
		delete(bill, item)
		return true
	}
	// Otherwise, reduce the quantity of the item and return true.
	bill[item] = newVal
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	itemExistsInBill := bill[item]
	if itemExistsInBill == 0 {
		// Return 0 and false if the item is not in the bill.
		return 0, false
	} else {
		// Otherwise, return the quantity of the item in the bill and true
		return bill[item], true
	}
}
