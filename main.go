package main

import (
	"fmt"
	"log"
	"net/http"

	"notes-crud-api/handlers"
	"notes-crud-api/store"

	"github.com/gorilla/mux"
)

const (
	port = ":8080"
)

func main() {
	// Initialize store
	notesStore := store.NewNotesStore()

	// Initialize handlers
	notesHandler := handlers.NewNotesHandler(notesStore)

	// Setup router
	router := mux.NewRouter()

	// API routes
	apiRouter := router.PathPrefix("/notes").Subrouter()
	apiRouter.HandleFunc("", notesHandler.CreateNote).Methods("POST")
	apiRouter.HandleFunc("", notesHandler.GetAllNotes).Methods("GET")
	apiRouter.HandleFunc("/{id}", notesHandler.GetNoteByID).Methods("GET")
	apiRouter.HandleFunc("/{id}", notesHandler.UpdateNote).Methods("PUT")
	apiRouter.HandleFunc("/{id}", notesHandler.DeleteNote).Methods("DELETE")

	// Apply CORS middleware
	handler := handlers.RegisterRoutes(router)

	// Print startup information
	printStartupInfo()

	// Start server
	log.Fatal(http.ListenAndServe(port, handler))
}

func printStartupInfo() {
	fmt.Println("Starting Notes CRUD API server...")
	fmt.Printf("Server is running on http://localhost%s\n", port)
	fmt.Println("\nAPI Endpoints:")
	fmt.Println("  POST   /notes       - Create a new note")
	fmt.Println("  GET    /notes       - Get all notes")
	fmt.Println("  GET    /notes/:id   - Get a specific note")
	fmt.Println("  PUT    /notes/:id   - Update a note")
	fmt.Println("  DELETE /notes/:id   - Delete a note")
	fmt.Println("\nPress Ctrl+C to stop the server")
}
