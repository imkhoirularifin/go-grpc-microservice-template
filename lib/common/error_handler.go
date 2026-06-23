package common

import (
	"github.com/gofiber/fiber/v2"
	"github.com/imkhoirularifin/go-grpc-microservice-template/lib/dto"
	"github.com/rs/zerolog/log"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	message := "internal server error"

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
		message = e.Message
	}

	log.Error().
		Err(err).
		Str("method", c.Method()).
		Str("path", c.Path()).
		Int("status", code).
		Msg("request failed")

	return c.Status(code).JSON(dto.ErrorResponse{
		Message: message,
	})
}

func NotFoundHandler(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
		Message: "route not found",
		Code:    "NOT_FOUND",
	})
}
