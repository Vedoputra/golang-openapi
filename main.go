//	@title			Golang API CRUD - Siswa
//	@version		1.0
//	@description	API sederhana CRUD Siswa (dummy/in-memory) + Swagger (Swaggo).
//	@host			localhost:3000
//	@BasePath		/api/v1
//	@schemes		http
package main

import (
	"log"

	"golang-api-crud/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/cors"
	swagger "github.com/swaggo/fiber-swagger"

	_ "golang-api-crud/docs" // penting: ini hasil generate swag init
)

func main() {
	app := fiber.New(fiber.Config{
		AppName: "Golang API CRUD - Siswa",
	})

	// Logging
	app.Use(logger.New())

	// CORS (penting untuk Flutter Web)
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Swagger UI: http://localhost:3000/swagger/index.html
	app.Get("/swagger/*", swagger.WrapHandler)

	// Routes API
	routes.SetupRoutes(app)

	// Redirect root ke swagger
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/swagger/index.html")
	})

	log.Println("Server running at http://localhost:3000")
	log.Fatal(app.Listen(":3000"))
}