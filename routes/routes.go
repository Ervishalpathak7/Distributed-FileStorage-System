package routes

import (
	"distributed-storage-system/handlers"

	"github.com/gofiber/fiber/v2"
)



func SetupRoutes(app *fiber.App) {

	app.Get("/api/v1/download", handlers.DownloadFileHandler)
	app.Post("/api/v1/upload", handlers.UploadFileHandler)
	
}