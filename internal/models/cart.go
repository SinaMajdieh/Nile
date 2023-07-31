package models

import "time"

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
	Items   []Item    `json:"items"`
	Expires time.Time 
}
