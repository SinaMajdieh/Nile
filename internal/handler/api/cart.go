package api

import (
	"Nile/internal/messages"
	"Nile/internal/models"
	"encoding/json"
	"time"

	"github.com/gofiber/fiber/v2"
)

// Adds an item to the sopping cart
func AddToCart(c *fiber.Ctx) error {

	// Creating a new empty item model
	item := new(models.Item)

	// Parsing the body request to the item model
	// Catch errors
	if err := c.BodyParser(item); nil != err {
		logger.Println(messages.BodyParser)
		return generateStatus(c, false, messages.BodyParser)
	}

	// Get existing cart
	// Catch errors
	cart, err := getCart(c)

	if nil != err {
		logger.Println(err)
		return generateStatus(c, false, err.Error())
	}

	// Add expiration date for cart cookie
	cart.Expires = time.Now().AddDate(0, 3, 0)

	// Item is not already in the cart
	item_index := cart.LookupItem(item.ProductId)
	if item_index == -1 {
		// Add it to the cart
		cart.PutItem(item)
	} else {
		cart.AddMoreOfItem(item_index, item.Quantity)
	}

	// Update the caret cookie
	setCart(c, cart)
	return generateStatus(c, true, "added item")

}

// Removes item from cart if exists
func RemoveFromCart(c *fiber.Ctx) error {

	// Creating a new empty item model
	item := new(models.Item)

	// Parsing the body request to the item model
	// Catch errors
	if err := c.BodyParser(item); nil != err {
		logger.Println(messages.BodyParser)
		return generateStatus(c, false, messages.BodyParser)
	}

	// Get existing cart
	// Catch errors
	cart, err := getCart(c)
	if nil != err {
		logger.Println(err)
		return generateStatus(c, false, err.Error())
	}

	// Add expiration date for cart cookie
	cart.Expires = time.Now().AddDate(0, 3, 0)

	// Lookup for the item in the cart
	item_index := cart.LookupItem(item.ProductId)
	if item_index == -1 {
		logger.Println(messages.ItemNotInCart)
		return generateStatus(c, false, messages.ItemNotInCart)
	}

	// Remove item
	if cart.RemoveItem(item_index, item.Quantity) {
		return generateStatus(c, true, "Removed")
	}

	return generateStatus(c, false, messages.ItemRemovalError)
}

// returns the user shopping cart
func GetMyCart(c *fiber.Ctx) error {

	// Get shopping cart
	// Catch errors
	cart, err := getCart(c)
	if err != nil {
		logger.Println(err.Error())
		return generateStatus(c, false, err.Error())
	}

	// Convert cart to json
	// Catch errors
	jsonByte, err := json.Marshal(cart)
	if nil != err {
		logger.Println("Unmarshal problem")
		return generateStatus(c, false, "Unmarshal problem")
	}

	// Convert jsonByte to string and send back
	jsonString := string(jsonByte)
	return generateStatus(c, true, jsonString)
}

func EmptyCart(c *fiber.Ctx) error {

	// Expire the cart
	cart := new(models.Cart)
	cart.Expires = time.Now().Add(-3 * time.Second)

	setCart(c, cart)
	return generateStatus(c, true, "cleared your cart")
}
