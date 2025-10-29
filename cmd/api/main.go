package main

import (
	"log"

	"github.com/gofiber/fiber"
)

func main() {
	// 1. INICIALIZACIÓN DE LA INFRAESTRUCTURA (MongoDB)
	// Suponiendo que tienes una función para conectar a Mongo
	// client, err := ConnectToMongoDB()
	// if err != nil {
	//     log.Fatal("Could not connect to MongoDB:", err)
	// }

	// 2. CONSTRUCCIÓN DE LAS CAPAS (De afuera hacia el centro)

	// Plataforma (Implementación de la DB)
	// userRepo := mongodb.NewUserRepository(client, "my_db", "users")

	// // Core/Servicios (Usa la interfaz de Repo)
	// userService := services.NewUserService(userRepo)

	// // Handlers/Controladores (Usa el Servicio)
	// userHandler := handlers.NewUserHandler(userService)

	// 3. CONFIGURACIÓN DE FIBER
	app := fiber.New()

	// Rutas (Conecta el Handler con Fiber)
	//routes.UserRoutes(app, userHandler)

	// 4. INICIO DEL SERVIDOR
	log.Fatal(app.Listen(":3000"))
}
