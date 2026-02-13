package cards

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	withData := []int{2,6,9}
	return withData
}


// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	// Source - https://stackoverflow.com/a/27252199␍
// Posted by laurent␍
// Retrieved 2026-02-12, License - CC BY-SA 3.0␍
//␍
	if index < len(slice) && index > -1 {
		return slice[index]
	} else {
	  return -1
	}
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if index < len(slice) && index > -1 {
		slice[index] = value
		return slice
	} else {
		slice = append(slice, value)
		return slice
	}
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	var new_slice []int
	new_slice = append(new_slice, values...)
	new_slice = append(new_slice, slice...)
	return new_slice
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if index < len(slice) && index > -1 {
		newSlice := append(slice[:index], slice[index+1:]...)
		return newSlice
	} else {
		return slice
	}

}
