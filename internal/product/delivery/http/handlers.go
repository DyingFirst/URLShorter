package http

import (
	"URLShorter/internal/models"
	"URLShorter/internal/product"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type Shorter struct {
	logger *logrus.Logger
	uc     product.UseCase
}

func NewDelivery(logger *logrus.Logger, uc product.UseCase) product.Delivery {
	return &Shorter{logger: logger, uc: uc}
}

func (d *Shorter) GetOriginalURL() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ShortedURL := c.Get("ShortedURL")
		OriginalURL, err := d.uc.GetOriginalURL(ShortedURL)
		if OriginalURL == "" {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		if err != nil {
			return c.SendStatus(fiber.StatusNotFound)
		}
		return c.SendString(OriginalURL)
	}
}

func (d *Shorter) NewShortURL() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var OriginalURL models.RequestOriginalURL
		if err := c.BodyParser(&OriginalURL); err != nil {
			d.logger.Error("Can't unmarshal URL")
			return c.SendStatus(fiber.StatusBadRequest)
		}
		ShortedURL, err := d.uc.NewShort(OriginalURL.OriginalURL)
		if ShortedURL == "" {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		if err != nil {

			return c.SendStatus(fiber.StatusBadRequest)
		}
		return c.SendString(ShortedURL)
	}
}
