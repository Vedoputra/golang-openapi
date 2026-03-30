package routes

import (
	"golang-api-crud/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")

	siswa := api.Group("/siswa")
	siswa.Get("/", handlers.GetAllSiswa)
	siswa.Get("/:id", handlers.GetSiswaByID)
	siswa.Post("/", handlers.CreateSiswa)
	siswa.Put("/:id", handlers.UpdateSiswa)
	siswa.Delete("/:id", handlers.DeleteSiswa)
}