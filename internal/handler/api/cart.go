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

	logger.Println(*item)

	// Get existing cart
	// Catch errors
	cart, err := getCart(c)
	if nil != err {
		logger.Println(err)
		return generateStatus(c, false, err.Error())
	}

	// Ensure cart exists
	if nil == cart {
		// Add item to the newly created cart model
		cart = new(models.Cart)
		cart.Items = append(cart.Items, *item)
		setCart(c, cart)
		return generateStatus(c, true, "added to the empty cart")
	}

	// Add expiration date for cart cookie
	cart.Expires = time.Now().AddDate(0, 3, 0)

	// item already exists
	for index, cartItem := range cart.Items {
		if cartItem.ProductId == item.ProductId {
			// Add more of the same product
			cart.Items[index].Quantity += item.Quantity
			setCart(c, cart)
			//return
			return generateStatus(c, true, "added to the previous ones")
		}
	}

	// Add item to the cart
	cart.Items = append(cart.Items, *item)
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

	logger.Println(*item)

	// Get existing cart
	// Catch errors
	cart, err := getCart(c)
	if nil != err {
		logger.Println(err)
		return generateStatus(c, false, err.Error())
	}

	// Ensure cart exists
	if nil == cart {
		// cart is empty
		// there is nothing to remove
		logger.Println(messages.EmptyCart)
		return generateStatus(c, false, messages.EmptyCart)
	}

	// Add expiration date for cart cookie
	cart.Expires = time.Now().AddDate(0, 3, 0)

	// item is in cart
	for index, cartItem := range cart.Items {
		if cartItem.ProductId == item.ProductId {
			// Ensure there are enough of them to remove
			if cart.Items[index].Quantity < item.Quantity {
				return generateStatus(c, false, "cant remove you dont have enough")
			} else if cart.Items[index].Quantity == item.Quantity {
				cart.Items[index] = cart.Items[len(cart.Items)-1] // Copy last element to index i.
				cart.Items[len(cart.Items)-1] = *new(models.Item) // Erase last element (write zero value).
				cart.Items = cart.Items[:len(cart.Items)-1]       // Truncate slice.
			} else {
				cart.Items[index].Quantity -= item.Quantity
			}
			setCart(c, cart)
			//return
			return generateStatus(c, true, "removed")
		}
	}

	return generateStatus(c, false, "item does not exist")
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

	// Cart does not exists
	if nil == cart {
		logger.Println("empty shopping cart")
		return generateStatus(c, false, "empty")
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
