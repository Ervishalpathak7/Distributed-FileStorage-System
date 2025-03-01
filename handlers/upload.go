package handlers

import (
	"io"
	"github.com/google/uuid"
	"distributed-storage-system/services"

	"github.com/gofiber/fiber/v2"
)


func UploadFileHandler(c *fiber.Ctx) error {

	// Extract metadata (assuming it's sent in headers or form fields)
	fileName := c.FormValue("fileName")
	fileType := c.FormValue("fileType")
	fileSize := c.FormValue("fileSize")

	if fileName == "" || fileType == "" || fileSize == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Missing file metadata"})
	}

	// Generate a unique file ID
	fileID := uuid.New().String()

	// File chunk size
	chunkSize := 1024 * 1024 
	buffer := make([]byte, chunkSize)
	chunkIndex := 0
	fileReader := c.Request().BodyStream()

	for {
		// Read a chunk of the file
		n, err := fileReader.Read(buffer)
		if err != nil && err != io.EOF {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error reading file"})
		}

		if n == 0 {
			break
		}

		// Upload the chunk to the storage service
		chunkId , err := services.UploadChunk(buffer[:n])
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error uploading file"})
		}
		chunkIndex++

		// Save the chunk data to the database
		err = services.SaveChunk(fileID, chunkIndex, chunkId)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error saving chunk"})
		}
	}

	// Save the file metadata to the database
	err := services.SaveFile(fileID, fileName, fileType, fileSize)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error saving file metadata"})
	}	

	return c.JSON(fiber.Map{"message": "File uploaded successfully", "fileID": fileID})
	
}








	