# Notes CRUD API

A simple CRUD (Create, Read, Update, Delete) API for managing notes, built with Go.

## Project Structure

```
.
├── main.go                 # Application entry point and server setup
├── models/
│   └── note.go            # Note model and DTOs (Data Transfer Objects)
├── store/
│   └── notes_store.go     # In-memory storage layer
├── handlers/
│   └── notes_handler.go   # HTTP request handlers
├── middleware/
│   └── cors.go            # CORS middleware
├── utils/
│   └── response.go        # HTTP response utilities
├── go.mod                 # Go module dependencies
└── README.md              # Project documentation
```

## Architecture

- **models**: Defines data structures (Note, request/response DTOs)
- **store**: Handles data persistence (in-memory storage with thread-safe operations)
- **handlers**: Processes HTTP requests and delegates to store layer
- **middleware**: HTTP middleware (CORS support)
- **utils**: Shared utility functions (JSON response helpers)

## Features

- **POST /notes** - Create a new note
- **GET /notes** - Retrieve all notes
- **GET /notes/:id** - Retrieve a specific note by ID
- **PUT /notes/:id** - Update a note
- **DELETE /notes/:id** - Delete a note

## Requirements

- Go 1.21 or higher

## Setup

1. Install dependencies:

```bash
go mod tidy
```

2. Run the server:

```bash
go run main.go
```

Or build and run:

```bash
go build -o notes-api .
./notes-api
```

The server will start on `http://localhost:8080`

## Usage

### API Endpoints

#### Create a Note

```bash
curl -X POST http://localhost:8080/notes \
  -H "Content-Type: application/json" \
  -d '{"title":"My Note","content":"This is the content"}'
```

#### Get All Notes

```bash
curl http://localhost:8080/notes
```

#### Get a Specific Note

```bash
curl http://localhost:8080/notes/1
```

#### Update a Note

```bash
curl -X PUT http://localhost:8080/notes/1 \
  -H "Content-Type: application/json" \
  -d '{"title":"Updated Note","content":"Updated content"}'
```

#### Delete a Note

```bash
curl -X DELETE http://localhost:8080/notes/1
```

## Notes

- Data is stored in-memory (will be lost when server restarts)
- All operations are logged to the console
- The API uses JSON for request/response format
- CORS is enabled for cross-origin requests
- Thread-safe operations using mutex locks
