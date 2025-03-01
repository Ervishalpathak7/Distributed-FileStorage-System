package models


// Chunk represents a chunk of a file
type Chunk struct {
	ID     string `json:"id"`
	Index  int    `json:"index"`
	FileID string `json:"file_id"`
}
