package models



// File represents a file in the system
type File struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Size     int64  `json:"size"`
	Chunks   []Chunk `json:"chunks"`
}
