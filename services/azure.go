package services

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)



var client *azblob.Client

// InitAzureClient initializes the Azure Blob Storage client
func InitAzureClient() (*azblob.Client, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	accountName := os.Getenv("AZURE_ACCOUNT_NAME")
	accountKey := os.Getenv("AZURE_ACCOUNT_KEY")
	
	// Create a connection string
	connectionString := fmt.Sprintf("DefaultEndpointsProtocol=https;AccountName=%s;AccountKey=%s;EndpointSuffix=core.windows.net", accountName, accountKey)

	client, err := azblob.NewClientFromConnectionString(connectionString, nil)
	if err != nil {
		return nil, err
	}

	// log the message if the client is successfully created
	log.Println("Azure Blob Storage client created")

	return client, nil
}


func UploadChunk(chunkData []byte) (string, error) {
	containerName := os.Getenv("AZURE_CONTAINER_NAME")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	blobId := uuid.New().String()

	// Upload the chunk to the storage service
	_, err := client.UploadBuffer(ctx, containerName, blobId, chunkData, nil)
	if err != nil {
		return "" , err
	}

	log.Printf("Uploaded chunk: %s\n", blobId)
	return blobId, nil
}


// DownloadChunk downloads a chunk from Azure Blob Storage
func DownloadChunk(chunkID string) ([]byte, error) {
	containerName := os.Getenv("AZURE_CONTAINER_NAME")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := client.DownloadStream(ctx, containerName, chunkID, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the chunk into memory
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}