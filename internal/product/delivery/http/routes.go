package http

import (
	"URLShorter/internal/product"
	"github.com/gofiber/fiber/v2"
)

func MapRoutes(client *fiber.App, d product.Delivery) {
	client.Get("/GetOriginalURL", d.GetOriginalURL())
	client.Post("/ShortURL", d.NewShortURL())
}
