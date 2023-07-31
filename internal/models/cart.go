package models

import (
	"time"
)

// Shopping cart models

// Items model in the shopping cart
// Including product id and quantity of it
type Item struct {
	ProductId int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

// Cart Model
// A list of Item models
type Cart struct {
	Items   []Item `json:"items"`
	Expires time.Time
}

// Looks up an item in the cart.
// Returning it's index if it existed.
// Returning -1 if the item was not in the cart
func (c Cart) LookupItem(id int) int {
	for i := range c.Items {
		if c.Items[i].ProductId == id {
			return i
		}
	}

	// item was not in the cart
	return -1
}

// Adds an  item to carts items.
// Do not care if it is already in the cart
func (c *Cart) PutItem(item *Item) {
	c.Items = append(c.Items, *item)
}

// Adds to the quantity of an item in the cart.
// Return true if added successfully.
// Returns false if item didn't exist.
func (c *Cart) AddMoreOfItem(index, count int) bool {
	// Ensure item exists
	if index >= len(c.Items) {
		return false
	}

	c.Items[index].Quantity += count
	return true
}

// Checks an item quantity.
// If it was more than count returns true.
// Otherwise returns false
func (c Cart) HasMore(index, count int) bool {
	// Ensure item exists
	if index > len(c.Items) {
		return false
	}

	// Has more than or equal to count
	if c.Items[index].Quantity >= count {
		return true
	}

	return false
}

func (c Cart) HasEqual(index, count int) bool {
	// Ensure item exists
	if index > len(c.Items) {
		return false
	}

	// Has more than or equal to count
	if c.Items[index].Quantity == count {
		return true
	}

	return false
}

// pops an item from the cart
func (c *Cart) PopItem(index int) {
	copy(c.Items[index:], c.Items[index+1:])
	c.Items[len(c.Items)-1] = *new(Item)
	c.Items = c.Items[:len(c.Items)-1]
}

// Removes Item from the cart.
// If item could be removed returns true and removes it.
// Otherwise returns false
func (c *Cart) RemoveItem(index, count int) bool {
	if c.HasMore(index, count) {
		c.Items[index].Quantity -= count
		return true
	} else if c.HasEqual(index, count) {
		c.PopItem(index)
		return true
	} else {
		// Has less quantity
		return false
	}
}
