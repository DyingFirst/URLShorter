package product

import "github.com/gofiber/fiber/v2"

type Delivery interface {
	GetOriginalURL() fiber.Handler
	NewShortURL() fiber.Handler
}
