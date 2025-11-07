// Lógica de la capa de transporte (Controladores/Controllers de Fiber)//contribuyentes

package handlers

import (
	"github.com/geovannymero/contribuyente-app/internal/core/services"
	"github.com/gofiber/fiber/v2"
)

type ContribuyenteHandler struct {
	contribuyenteService *services.ContribuyenteService
}

func NewContribuyenteHandler(us *services.ContribuyenteService) *ContribuyenteHandler {
	return &ContribuyenteHandler{
		contribuyenteService: us,
	}
}

func (h *ContribuyenteHandler) Get(c *fiber.Ctx) error {

	// 1. Obtener y validar parámetros de paginación de la URL
	page := c.QueryInt("page", 1)    // Valor por defecto: 1
	limit := c.QueryInt("limit", 10) // Valor por defecto: 10

	contri, err := h.contribuyenteService.Get("dasd", page, limit)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": ""})
	}

	return c.Status(fiber.StatusOK).JSON(contri)
	//return response.Ok(contri)
}
