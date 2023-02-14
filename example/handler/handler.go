package handler

import "github.com/gofiber/fiber/v2"

var (
	mockResponseJSON = map[string]interface{}{
		"key":          "value",
		"nickname":     "wat",
		"lucky_number": 11,
		"citizen_id":   "1234567890121", // <<< sensitive value
	}
)

func OriginalHandler(c *fiber.Ctx) error {
	return c.JSON(mockResponseJSON)
}

func StringHandler(c *fiber.Ctx) error {
	return c.SendString("a simple string, can be return with no any error.")
}

func SensitiveHandler(c *fiber.Ctx) error {
	return c.JSON(mockResponseJSON)
}
