package services

import (
	"context"
	"distributed-storage-system/models"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)


var conn *pgx.Conn


// InitPostgresClient initializes the Postgres client
func InitPostgresClient() (*pgx.Conn, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	
	context , cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGDATABASE"),
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGSSLMODE"),
	)

	conn, err := pgx.Connect(context, connStr)
	if err != nil {
		return nil, err
	}

	// log the message if the client is successfully created
	log.Println("Postgres client created")

	return conn, nil
}


// SaveChunk saves the chunk data to the database
func SaveChunk(fileID string, chunkIndex int, chunkID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := conn.Exec(ctx, "INSERT INTO chunks (file_id, chunk_index, chunk_id) VALUES ($1, $2, $3)", fileID, chunkIndex, chunkID)

	if err != nil {
		return err
	}

	log.Printf("Saved chunk: %s\n", chunkID)
	return nil
}

// SaveFile saves the file metadata to the database
func SaveFile(fileID string, fileName string, fileType string, fileSize string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := conn.Exec(ctx, "INSERT INTO files (file_id, file_name, file_type, file_size) VALUES ($1, $2, $3, $4)", fileID, fileName, fileType, fileSize)

	if err != nil {
		return err
	}

	log.Printf("Saved file: %s\n", fileID)
	return nil
}

// GetFileMetadata retrieves file metadata from the database
func GetFileMetadata(fileID string) (*models.File, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	file := &models.File{}
	err := conn.QueryRow(ctx, "SELECT file_name, file_type FROM files WHERE file_id = $1", fileID).
		Scan(&file.Name, &file.Type)

	if err != nil {
		return nil, err
	}

	return file, nil
}

// GetFileChunks retrieves all chunk IDs for a given file
func GetFileChunks(fileID string) ([]models.Chunk, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	rows, err := conn.Query(ctx, "SELECT chunk_id FROM chunks WHERE file_id = $1 ORDER BY chunk_index", fileID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var chunks []models.Chunk
	for rows.Next() {
		var chunk models.Chunk
		if err := rows.Scan(&chunk.ID); err != nil {
			return nil, err
		}
		chunks = append(chunks, chunk)
	}

	return chunks, nil
}