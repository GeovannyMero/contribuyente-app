package main

import (
	"fmt"
	"log"

	"github.com/geovannymero/contribuyente-app/internal/core/services"
	"github.com/geovannymero/contribuyente-app/internal/handlers"
	"github.com/geovannymero/contribuyente-app/internal/infraestructure/config"
	"github.com/geovannymero/contribuyente-app/internal/infraestructure/mongodb"
	"github.com/geovannymero/contribuyente-app/internal/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// 1. INICIALIZACIÓN DE LA INFRAESTRUCTURA (MongoDB)
	// Suponiendo que tienes una función para conectar a Mongo
	client, err := config.ConnectToMongoDB()
	if err != nil {
		log.Fatal("Could not connect to MongoDB:", err)
	}

	// 2. CONSTRUCCIÓN DE LAS CAPAS (De afuera hacia el centro)

	// Plataforma (Implementación de la DB)
	// userRepo := mongodb.NewUserRepository(client, "my_db", "users")
	contriRepo := mongodb.NewMongoRepository(client, "testdb", "test")

	// // Core/Servicios (Usa la interfaz de Repo)
	// userService := services.NewUserService(userRepo)
	contriservice := services.NewContribuyenteService(contriRepo)
	// // Handlers/Controladores (Usa el Servicio)
	userHandler := handlers.NewContribuyenteHandler(contriservice)

	// 3. CONFIGURACIÓN DE FIBER
	app := fiber.New()

	app.Use(cors.New())

	// Rutas (Conecta el Handler con Fiber)
	routes.UserRoutes(app, userHandler)

	routes.InitStaticRoutes(app)

	fmt.Println("Iniciando en el puerto 3000")
	// 4. INICIO DEL SERVIDOR
	log.Fatal(app.Listen(":3000"))
}
