package handlers

import (
	"distributed-storage-system/services"
	"fmt"

	"github.com/gofiber/fiber/v2"
)



func DownloadFileHandler(c *fiber.Ctx) error {
	fileID := c.Params("fileID")
	if fileID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Missing file ID"})
	}

	// Fetch file metadata
	file, err := services.GetFileMetadata(fileID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "File not found"})
	}

	// Fetch chunks from the database
	chunks, err := services.GetFileChunks(fileID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error fetching file chunks"})
	}

	// Set response headers for file download
	c.Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", file.Name))
	c.Set("Content-Type", file.Type)

	// Stream the chunks back to the client
	for _, chunk := range chunks {
		chunkData, err := services.DownloadChunk(chunk.ID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error downloading chunk"})
		}
		if _, err := c.Write(chunkData); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error streaming chunk"})
		}
	}

	return nil
}