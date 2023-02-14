package sensitive

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func New(config ...Config) fiber.Handler {
	// set default config
	cfg := configDefault(config...)
	logger := NewLogger(cfg.DebugMode)
	logger.Printf("[Sensitive] Config: %+v\n", cfg)

	return func(c *fiber.Ctx) error {
		// Don't execute middleware if Next returns true
		if cfg.Next != nil && cfg.Next(c) {
			return c.Next()
		}

		// Let request to continue execute
		c.Next()

		// Intercept the response body
		interceptBody := c.Response().Body()
		logger.Print("[Sensitive] Original Body:", string(interceptBody))

		var mapResponseBody fiber.Map
		err := json.Unmarshal(interceptBody, &mapResponseBody)
		if err != nil {
			return c.SendString(string(interceptBody))
		}

		for _, key := range cfg.Keys {
			switch mapResponseBody[key].(type) {
			case string:
				break
			default:
				continue
			}
			target := mapResponseBody[key].(string)
			length := len(target)
			if length > 2 {
				firstChar := target[:1]
				middle := strings.Repeat(cfg.Mark, len(target[1:length-1]))
				lastChar := target[length-1:]

				logger.Print("[Sensitive] > before", key, "=", mapResponseBody[key])
				mapResponseBody[key] = fmt.Sprintf("%s%s%s",
					firstChar,
					middle,
					lastChar,
				)
				logger.Print("[Sensitive] > after ", key, "=", mapResponseBody[key])
			}
		}

		logger.Print("[Sensitive] Blinded Body:", mapResponseBody)
		return c.JSON(mapResponseBody)
	}
}
