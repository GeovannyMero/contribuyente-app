package main

import (
	"context"
	"fmt"

	"github.com/geovannymero/contribuyente-app/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	fmt.Println("iniciando")

	app := fiber.New()

	app.Use(cors.New())

	//mongoDB
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017/testdb"))

	if err != nil {
		panic(err)
	}

	app.Static("/", "./client/dist")

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello, World!")
	// })

	// app.Get("/:value", func(c *fiber.Ctx) error {
	// 	return c.SendString("value: " + c.Params("value"))
	// })

	app.Get("/api/contribuyentes", func(c *fiber.Ctx) error {

		var province = c.Query("provincia")
		razon_social := c.Query("razon_social")

		filter := bson.D{} //filtros

		var contribuyentes []models.Contribuyente

		//Se valida para cada query params
		if province != "" {
			filter = append(filter, bson.E{Key: "codigo_juridiccion", Value: province})
		}

		if razon_social != "" {
			filter = append(filter, bson.E{Key: "razon_social", Value: razon_social})
		}

		// filter := bson.D{{Key: "codigo_juridiccion", Value: province}} //, {Key: "razon_social", Value: razon_social}

		coll := client.Database("testdb").Collection("test")
		opts := options.Find().SetLimit(20)
		result, err := coll.Find(context.TODO(), filter, opts)

		if err != nil {
			panic(err)

		}
		for result.Next(context.TODO()) {
			var con models.Contribuyente

			result.Decode(&con)
			contribuyentes = append(contribuyentes, con)

		}

		return c.JSON(&fiber.Map{
			"data": contribuyentes,
		})

	})

	app.Listen(":3000")

	fmt.Println("Iniciando en el puerto 3000")
}
