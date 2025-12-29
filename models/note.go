package models

import "time"

// Note represents a note in the system
type Note struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// CreateNoteRequest represents the request payload for creating a note
type CreateNoteRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// UpdateNoteRequest represents the request payload for updating a note
type UpdateNoteRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error string `json:"error"`
}

// DeleteResponse represents a successful delete response
type DeleteResponse struct {
	Message string `json:"message"`
}

type NoteResponse struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
}