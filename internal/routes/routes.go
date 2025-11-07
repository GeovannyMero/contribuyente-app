package routes

import (
	"github.com/geovannymero/contribuyente-app/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App, h *handlers.ContribuyenteHandler) {
	api := app.Group("/api/v1")
	api.Get("/contribuyentes", h.Get)

	// api.Static("/", "./client/dist")
}

// InitStaticRoutes configura el enrutamiento para archivos estáticos.
func InitStaticRoutes(app *fiber.App) {
	// fiber.Static configura Fiber para servir archivos desde una ruta del sistema.

	// 1. Ruta URL: /static
	//    Cualquier solicitud que comience con /static (ej: /static/imagen.jpg)

	// 2. Directorio Local: ./web/static
	//    Fiber buscará el archivo solicitado en esta carpeta.

	app.Static("/", "../../client/dist")

	// Opcional: Si quieres que el directorio raíz de la API sirva tu frontend, usa "/"
	// app.Static("/", "./web/static/html")
}
