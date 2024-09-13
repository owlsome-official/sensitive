package main

import (
	"example/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/owlsome-official/sensitive"
)

func main() {
	app := fiber.New()

	originalGroup := app.Group("original")
	originalGroup.Get("/", handler.OriginalHandler)     // /original
	originalGroup.Get("/string", handler.StringHandler) // /original/string

	sensitiveGroup := app.Group("blinded")
	sensitiveGroup.Use(sensitive.New(sensitive.Config{
		Keys: []string{
			"citizen_id",
		},
		DebugMode: true,
	}))
	sensitiveGroup.Get("/", handler.SensitiveHandler) // /blinded

	port := "5000"
	app.Listen(":" + port)
}
