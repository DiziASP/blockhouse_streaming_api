package responses

import (
	"blockhouse_streaming_api/pkg/file/json"
	"github.com/gofiber/fiber/v2"
)

type General struct {
	StatusCode int         `json:"-"`
	Status     string      `json:"status"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

func (g *General) JSON(c *fiber.Ctx) error {
	return c.Status(g.StatusCode).JSON(g)
}

func BindingGeneral(data interface{}) General {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return DefaultErrorResponse
	}
	var response General
	if err := json.Unmarshal(jsonData, &response); err != nil {
		return DefaultErrorResponse
	}
	return response
}
