package routes

import (
	"github.com/geovannymero/contribuyente-app/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App, h *handlers.ContribuyenteHandler) {
	api := app.Group("/api/v1")
	api.Get("/contribuyentes", h.Get)
}
