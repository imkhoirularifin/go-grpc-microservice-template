package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/imkhoirularifin/go-grpc-microservice-template/lib/dto"
)

type HealthHandler struct{}

func NewHealthHandler(r fiber.Router) {
	h := &HealthHandler{}
	r.Get("/healthz", h.Healthz)
	r.Get("/readyz", h.Readyz)
}

func (h *HealthHandler) Healthz(c *fiber.Ctx) error {
	return c.JSON(dto.Response{Message: "ok"})
}

func (h *HealthHandler) Readyz(c *fiber.Ctx) error {
	return c.JSON(dto.Response{Message: "ready"})
}
