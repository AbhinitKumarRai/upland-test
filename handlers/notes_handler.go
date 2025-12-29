package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"notes-crud-api/models"
	"notes-crud-api/store"
	"notes-crud-api/utils"

	"github.com/gorilla/mux"
)

// NotesHandler handles HTTP requests for notes
type NotesHandler struct {
	store *store.NotesStore
}

// NewNotesHandler creates a new NotesHandler instance
func NewNotesHandler(store *store.NotesStore) *NotesHandler {
	return &NotesHandler{
		store: store,
	}
}

// CreateNote handles POST /notes - Create a new note
func (h *NotesHandler) CreateNote(w http.ResponseWriter, r *http.Request) {
	var req models.CreateNoteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if req.Title == "" || req.Content == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Title and content are mandatory")
		return
	}

	note := h.store.CreateNote(req.Title, req.Content)
	// fmt.Printf("[CREATE] Note created: ID=%d, Title=%s\n", note.ID, note.Title)
	utils.RespondWithJSON(w, http.StatusCreated, map[string]any{"NoteId": note.ID})
}

// GetAllNotes handles GET /notes - Retrieve all notes
func (h *NotesHandler) GetAllNotes(w http.ResponseWriter, r *http.Request) {
	notes := h.store.GetAllNotes()
	if len(notes) == 0 {
		utils.RespondWithError(w, http.StatusNotFound, "No notes found")
		return
	}
	// fmt.Printf("[GET ALL] Retrieved %d notes\n", len(notes))
	utils.RespondWithJSON(w, http.StatusOK, notes)
}

// GetNoteByID handles GET /notes/:id - Retrieve a specific note by ID
func (h *NotesHandler) GetNoteByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid note ID")
		return
	}

	note, exists := h.store.GetNoteByID(id)
	if !exists {
		utils.RespondWithError(w, http.StatusNotFound, "Note not found")
		return
	}

	fmt.Printf("[GET BY ID] Retrieved note: ID=%d, Title=%s\n", note.ID, note.Title)
	utils.RespondWithJSON(w, http.StatusOK, note)
}

// UpdateNote handles PUT /notes/:id - Update a note
func (h *NotesHandler) UpdateNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid note ID")
		return
	}

	var req models.UpdateNoteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if req.Title == "" || req.Content == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Title and content are required")
		return
	}

	note, exists := h.store.UpdateNote(id, req.Title, req.Content)
	if !exists {
		utils.RespondWithError(w, http.StatusNotFound, "Note not found")
		return
	}

	fmt.Printf("[UPDATE] Note updated: ID=%d, Title=%s\n", note.ID, note.Title)
	utils.RespondWithJSON(w, http.StatusOK, note)
}

// DeleteNote handles DELETE /notes/:id - Delete a note
func (h *NotesHandler) DeleteNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid note ID")
		return
	}

	exists := h.store.DeleteNote(id)
	if !exists {
		utils.RespondWithError(w, http.StatusNotFound, "Note not found")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, models.DeleteResponse{
		Message: "Note deleted successfully",
	})
}


// CORS middleware enables Cross-Origin Resource Sharing
func RegisterRoutes(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}