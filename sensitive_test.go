package sensitive

import (
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestSensitive(t *testing.T) {
	tests := []struct {
		name             string
		config           Config
		originalResponse fiber.Map
		blindedResponse  fiber.Map
	}{
		{
			name:   "empty string",
			config: Config{Keys: []string{"citizen_id"}},
			originalResponse: fiber.Map{
				"key":        "value",
				"citizen_id": "", // <<< sensitive value
			},
			blindedResponse: fiber.Map{
				"key":        "value",
				"citizen_id": "", // <<< sensitive value
			},
		},
		{
			name:   "no blind when have only 2 characters value",
			config: Config{Keys: []string{"onlytwo"}},
			originalResponse: fiber.Map{
				"k":       "v",
				"onlytwo": "az", // <<< sensitive value
			},
			blindedResponse: fiber.Map{
				"k":       "v",
				"onlytwo": "az", // <<< sensitive value
			},
		},
		{
			name:   "blinded x with 1 of 3 characters value",
			config: Config{Keys: []string{"somekey"}},
			originalResponse: fiber.Map{
				"k":       "v",
				"somekey": "123", // <<< sensitive value
			},
			blindedResponse: fiber.Map{
				"k":       "v",
				"somekey": "1x3", // <<< sensitive value
			},
		},
		{
			name:   "multi keys",
			config: Config{Keys: []string{"k1", "k2", "k3"}},
			originalResponse: fiber.Map{
				"static": "value",
				"k1":     "abcdefg",    // <<< sensitive value
				"k2":     "0987654321", // <<< sensitive value
				"k3":     "A0000000Z",  // <<< sensitive value
			},
			blindedResponse: fiber.Map{
				"static": "value",
				"k1":     "axxxxxg",    // <<< sensitive value
				"k2":     "0xxxxxxxx1", // <<< sensitive value
				"k3":     "AxxxxxxxZ",  // <<< sensitive value
			},
		},
		{
			name:   "blinded with 'T'",
			config: Config{Keys: []string{"key"}, Mark: "T"},
			originalResponse: fiber.Map{
				"key": "1234567890", // <<< sensitive value
			},
			blindedResponse: fiber.Map{
				"key": "1TTTTTTTT0", // <<< sensitive value
			},
		},
		{
			name:             "empty body",
			config:           Config{},
			originalResponse: fiber.Map{},
			blindedResponse:  fiber.Map{},
		},
		{
			name:   "not blind when target key is not type string (float)",
			config: Config{Keys: []string{"not_string"}},
			originalResponse: fiber.Map{
				"string":     "1234567890",
				"not_string": 1234.12, // <<< sensitive value
			},
			blindedResponse: fiber.Map{
				"string":     "1234567890",
				"not_string": 1234.12, // <<< sensitive value
			},
		},
		{
			name:   "not blind when target key is not type string (map[string]interface{})",
			config: Config{Keys: []string{"not_string"}},
			originalResponse: fiber.Map{
				"string":     "1234567890",
				"not_string": map[string]interface{}{"k": "v"}, // <<< sensitive value
			},
			blindedResponse: fiber.Map{
				"string":     "1234567890",
				"not_string": map[string]interface{}{"k": "v"}, // <<< sensitive value
			},
		},
		{
			name:             "empty body",
			config:           Config{},
			originalResponse: fiber.Map{},
			blindedResponse:  fiber.Map{},
		},
		{
			name: "call Next() once",
			config: Config{Next: func(c *fiber.Ctx) bool {
				return true
			}},
			originalResponse: nil,
			blindedResponse:  nil,
		},
	}
	route := "/_test"
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			app := fiber.New()
			app.Use(New(tt.config))
			app.Get(route, func(c *fiber.Ctx) error {
				return c.JSON(tt.originalResponse)
			})

			req := httptest.NewRequest("GET", route, nil)
			resp, _ := app.Test(req, -1)

			if resp.StatusCode == fiber.StatusOK {
				body, _ := io.ReadAll(resp.Body)
				var actualResponseBody fiber.Map
				_ = json.Unmarshal(body, &actualResponseBody)
				assert.Equal(t, tt.blindedResponse, actualResponseBody)
			}
		})
	}

	t.Run("Test error intercept body Unmarshal", func(t *testing.T) {
		originalResponse := "Some string"
		expected := "Some string"

		app := fiber.New()
		app.Use(New())
		app.Get(route, func(c *fiber.Ctx) error {
			return c.SendString(originalResponse)
		})

		req := httptest.NewRequest("GET", route, nil)
		resp, _ := app.Test(req, -1)
		body, _ := io.ReadAll(resp.Body)
		assert.Equal(t, string(body), expected)
	})
}
