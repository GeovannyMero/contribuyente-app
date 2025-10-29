// Lógica de la capa de transporte (Controladores/Controllers de Fiber)//contribuyentes

package handlers

import (
	"github.com/geovannymero/contribuyente-app/internal/core/services"
	"github.com/gofiber/fiber/v2"
)

type ContribuyenteHandler struct {
	contribuyenteService *services.ContribuyenteService
}

func (h *ContribuyenteHandler) Get(c *fiber.Ctx) error {

	contri, err := h.contribuyenteService.Get("dasd")

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": ""})
	}

	return c.Status(fiber.StatusOK).JSON(contri)
}
