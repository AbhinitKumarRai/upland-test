package store

import (
	"notes-crud-api/models"
	"sync"
	"time"
)

// NotesStore manages the in-memory storage of notes
type NotesStore struct {
	mu     sync.RWMutex
	notes  map[int]*models.Note
	nextID int
}

// NewNotesStore creates a new NotesStore instance
func NewNotesStore() *NotesStore {
	return &NotesStore{
		notes:  make(map[int]*models.Note),
		nextID: 1,
	}
}

// CreateNote adds a new note to the store
func (ns *NotesStore) CreateNote(title, content string) *models.NoteResponse {
	ns.mu.Lock()
	defer ns.mu.Unlock()

	note := &models.Note{
		ID:        ns.nextID,
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	ns.notes[note.ID] = note
	ns.nextID++
	return &models.NoteResponse{
		ID:        note.ID,
		Title:     note.Title,
		Content:   note.Content,
	}
}

// GetAllNotes returns all notes in the store
func (ns *NotesStore) GetAllNotes() []*models.NoteResponse {
	ns.mu.RLock()
	defer ns.mu.RUnlock()

	notes := make([]*models.NoteResponse, 0, len(ns.notes))
	for _, note := range ns.notes {
		notes = append(notes, &models.NoteResponse{
			ID:        note.ID,
			Title:     note.Title,
			Content:   note.Content,
		})
	}
	return notes
}

// GetNoteByID returns a note by its ID
func (ns *NotesStore) GetNoteByID(id int) (*models.NoteResponse, bool) {
	ns.mu.RLock()
	defer ns.mu.RUnlock()

	if note, exists := ns.notes[id]; exists {
		return &models.NoteResponse{
			ID:        note.ID,
			Title:     note.Title,
			Content:   note.Content,
		}, true
	}
	return nil, false
}

// UpdateNote updates an existing note
func (ns *NotesStore) UpdateNote(id int, title, content string) (*models.NoteResponse, bool) {
	ns.mu.Lock()
	defer ns.mu.Unlock()

	note, exists := ns.notes[id]
	if !exists {
		return nil, false
	}

	note.Title = title
	note.Content = content
	note.UpdatedAt = time.Now()
	return &models.NoteResponse{
		ID:        note.ID,
		Title:     note.Title,
		Content:   note.Content,
	}, true
}

// DeleteNote removes a note by its ID
func (ns *NotesStore) DeleteNote(id int) bool {
	ns.mu.Lock()
	defer ns.mu.Unlock()

	_, exists := ns.notes[id]
	if !exists {
		return false
	}

	delete(ns.notes, id)
	return true
}

