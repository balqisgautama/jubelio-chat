package util

import (
	"github.com/gofiber/fiber/v2"
)

func ReadBody(request *fiber.Request) (output string, bodySize int, err error) {
	return string(request.Body()), len(request.Body()), nil
}
