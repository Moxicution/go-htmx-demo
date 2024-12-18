package main

import (
	"errors"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// Note is the Note structure
type Note struct {
	ID          int       `json:"id,omitempty"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Rating      int       `json:"rating"`
	Completed   bool      `json:"completed,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
}

// ConvertedNote is the ConvertedNote structure
type ConvertedNote struct {
	ID          int    `json:"id,omitempty"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Rating      int    `json:"rating"`
	Completed   bool   `json:"completed,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
}

// CreateNote creates a Note
func (n *Note) CreateNote() (Note, error) {
	db := GetConnection()

	query := `INSERT INTO notes (title, description, rating, created_at) VALUES(?, ?, ?, ?) RETURNING *`

	stmt, err := db.Prepare(query)
	if err != nil {
		return Note{}, err
	}

	defer stmt.Close()

	var newNote Note
	err = stmt.QueryRow(
		n.Title,
		n.Description,
		n.Rating,
		time.Now().UTC(),
	).Scan(
		&newNote.ID,
		&newNote.Title,
		&newNote.Description,
		&newNote.Rating,
		&newNote.Completed,
		&newNote.CreatedAt,
	)
	if err != nil {
		return Note{}, err
	}

	/* if i, err := result.RowsAffected(); err != nil || i != 1 {
		return errors.New("error: an affected row was expected")
	} */

	return newNote, nil
}

// GetAllNotes gets all notes
func (n *Note) GetAllNotes(offset int) ([]Note, error) {
	db := GetConnection()
	query := fmt.Sprintf("SELECT * FROM notes ORDER BY created_at DESC LIMIT 5 OFFSET %d", offset)

	rows, err := db.Query(query)
	if err != nil {
		return []Note{}, err
	}
	// Cerramos el recurso
	defer rows.Close()

	notes := []Note{}
	for rows.Next() {
		err = rows.Scan(&n.ID, &n.Title, &n.Description, &n.Rating, &n.Completed, &n.CreatedAt)
		if err != nil {
			notes = append(notes, *n)
		}
	}

	return notes, nil
}

// GetNoteByID gets a Note by ID
func (n *Note) GetNoteByID() (Note, error) {
	db := GetConnection()

	query := `SELECT * FROM notes WHERE id=?`

	stmt, err := db.Prepare(query)
	if err != nil {
		return Note{}, err
	}

	defer stmt.Close()

	var recoveredNote Note
	err = stmt.QueryRow(
		n.ID,
	).Scan(
		&recoveredNote.ID,
		&recoveredNote.Title,
		&recoveredNote.Description,
		&recoveredNote.Rating,
		&recoveredNote.Completed,
		&recoveredNote.CreatedAt,
	)
	if err != nil {
		return Note{}, err
	}

	return recoveredNote, nil
}

// UpdateNote updates a Note
func (n *Note) UpdateNote() (Note, error) {
	db := GetConnection()

	query := `UPDATE notes SET completed=? WHERE id=? RETURNING *`

	stmt, err := db.Prepare(query)
	if err != nil {
		return Note{}, err
	}

	defer stmt.Close()

	var updatedNote Note
	err = stmt.QueryRow(
		!n.Completed,
		n.ID,
	).Scan(
		&updatedNote.ID,
		&updatedNote.Title,
		&updatedNote.Description,
		&updatedNote.Rating,
		&updatedNote.Completed,
		&updatedNote.CreatedAt,
	)
	if err != nil {
		return Note{}, err
	}

	return updatedNote, nil
}

// DeleteNote deletes a Note
func (n *Note) DeleteNote() error {
	db := GetConnection()

	query := `DELETE FROM notes WHERE id=?`

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(n.ID)
	if err != nil {
		return err
	}

	if i, err := result.RowsAffected(); err != nil || i != 1 {
		return errors.New("an affected row was expected")
	}

	return nil
}

func convertDateTime(note Note, timeZone string) ConvertedNote {
	loc, _ := time.LoadLocation(timeZone)
	convertedNote := ConvertedNote{
		ID:          note.ID,
		Title:       note.Title,
		Description: note.Description,
		Rating:      note.Rating,
		Completed:   note.Completed,
		CreatedAt:   note.CreatedAt.In(loc).Format(time.RFC822Z),
	}

	return convertedNote
}
