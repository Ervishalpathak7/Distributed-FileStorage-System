# Distributed File Storage System (Go + Azure + PostgreSQL)

A distributed file storage system built with Go, using Azure Blob Storage for chunk storage and PostgreSQL for metadata management. Files are uploaded in chunks, stored securely, and can be downloaded seamlessly.

## 📂 Features

- **Chunk-Based Uploads:** Files are split into chunks and stored individually in Azure Blob Storage.
- **Metadata Tracking:** File and chunk metadata (file ID, name, type, size, and chunk indices) are stored in PostgreSQL.
- **Streaming Downloads:** Chunks are streamed back to the client in sequence during downloads.
- **Memory Efficiency:** Handles large files without overloading server memory.

## 🛠️ Tech Stack

- **Go** (with Fiber framework)
- **Azure Blob Storage** (for chunk storage)
- **PostgreSQL** (for metadata)
- **Docker** (optional, for local dev)

## ⚙️ Folder Structure

```
Distributed File Storage
├── config                # Config files (DB, Azure)
├── db                   # Database migrations
├── handlers             # Route handlers
├── models               # Database models
├── routes               # API routes
├── services             # Core logic for file storage and retrieval
├── utils                # Utility functions
├── main.go              # Application entry point
└── Dockerfile           # Docker setup
```

## 🚀 Getting Started

### 1. Environment Setup

Set up the following environment variables:

```
AZURE_STORAGE_ACCOUNT_NAME=your_storage_account_name
AZURE_STORAGE_ACCOUNT_KEY=your_storage_account_key
AZURE_CONTAINER_NAME=file-chunks
DB_HOST=your_postgres_host
DB_USER=your_postgres_user
DB_PASSWORD=your_postgres_password
DB_NAME=your_postgres_db
DB_PORT=5432
```

### 2. Install Dependencies

```sh
go mod tidy
```

### 3. Run Migrations

```sh
migrate -path db/migrations -database "$DATABASE_URL" up
```

### 4. Start the Server

```sh
go run main.go
```

The app will run at: `http://localhost:3000`

## 🔑 API Endpoints

### Upload File

```
POST /upload
```

- **Headers/Form Data:** `fileName`, `fileType`, `fileSize`
- **Body:** File stream (sent in chunks)

### Download File

```
GET /download/:fileID
```

- **Params:** `fileID`

The server will stream the file back to the client.

## 🚀 Deployment

You can deploy the app using Docker or any cloud platform that supports Go apps.

Example Docker build:

```sh
docker build -t distributed-file-storage .
```

## 🧠 Future Enhancements

- **Replication & Redundancy:** Store chunks across multiple containers.
- **Authentication & Authorization:** Add secure access control.
- **File Integrity Checks:** Verify chunk integrity with checksums.

---

Ready to showcase this on GitHub? Let me know if you want any changes! 🚀

