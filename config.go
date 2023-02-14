package sensitive

import (
	"github.com/gofiber/fiber/v2"
)

type Config struct {
	// Optional. Default: nil
	Next func(c *fiber.Ctx) bool

	// Required. Default: []
	Keys []string

	// Optional. Default: "x"
	Mark string

	// Optional. Default: false
	DebugMode bool
}

var ConfigDefault = Config{
	Next:      nil,
	Keys:      []string{},
	Mark:      "x",
	DebugMode: false,
}

func configDefault(config ...Config) Config {
	// Return default config if nothing provided
	if len(config) < 1 {
		return ConfigDefault
	}

	// Override default config
	cfg := config[0]

	// set default values
	if cfg.Keys == nil {
		cfg.Keys = ConfigDefault.Keys
	}

	if cfg.Mark == "" {
		cfg.Mark = ConfigDefault.Mark
	}

	return cfg
}
