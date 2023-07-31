package api

import (
	"Nile/internal/models"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

const (
	// Name of the cart cookie
	cartCookieName = "cart"
)

// Sets the cart cookie for the user
// Marshalling the cart model into json and storing it as a cookie
func setCart(c *fiber.Ctx, cart *models.Cart) error {

	// Marshalling the cart model into json bytes
	// catching the error
	jsonByte, err := json.Marshal(cart)
	if nil != err {
		logger.Println(err)
		return err
	}

	// Converting the json bytes into string
	jsonString := string(jsonByte)

	//TEMP: logging the result string
	logger.Println(jsonString)

	// Setting the shopping cart cookie
	c.Cookie(&fiber.Cookie{
		Name:    cartCookieName,
		Value:   jsonString,
		Expires: cart.Expires,
	})

	return nil
}

// getCart checks if there was an existing shopping cart cookie.
// Returns a cart model and an error result.
// Returns nil as the cart model and the error if there was one.
// Returns the cart model and nil as the error if there was no errors.
func getCart(c *fiber.Ctx) (*models.Cart, error) {

	// String value of the shopping cart
	jsonString := c.Cookies(cartCookieName)

	// Ensure cookie exists
	if jsonString == "" {
		logger.Println("No shopping cart")
		return nil, nil
	}

	// Create a new empty cart model
	cart := new(models.Cart)

	// Unmarshal the json value
	// catch errors
	err := json.Unmarshal([]byte(jsonString), cart)
	if nil != err {
		logger.Println("Couldn't unmarshal")
		return nil, err
	}

	return cart, nil

}
