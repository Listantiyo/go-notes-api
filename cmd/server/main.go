package main

import (
	"fmt"
	"log/slog"
	"net/http"

	"database/sql"
	"encoding/json"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	_ "modernc.org/sqlite"
)

type NewNotesRequest struct {
	Title string		`json:"title"`
	Descriptions string `json:"descriptions"`
}

type Notes struct {
	ID string 		`json:"id" db:"id"`
	Title string		`json:"title" db:"title"`
	Descriptions string	`json:"descriptions" db:"descriptions"`
}

func prepareDB(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS notes (
		id TEXT PRIMARY KEY,
		title TEXT,
		descriptions TEXT
	)`

	
	if _, err := db.Exec(query); err != nil {
		return err
	}

	return nil
}

func main() {
	db, err := sql.Open("sqlite", "./database/database.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	r := chi.NewRouter()

	if err := prepareDB(db); err != nil {
		slog.Error("Failed ini db", "error", err)
		panic("failed to init database")
	}


	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Notes API 🚀")
	})
	

	r.Post("/note", func(w http.ResponseWriter, r *http.Request) {
		var req NewNotesRequest
		json.NewDecoder(r.Body).Decode(&req)

		query := `INSERT INTO notes (id, title, descriptions) VALUES (?, ?, ?)`
		id := uuid.New().String()

		notes := Notes{
			ID: id,
			Title: req.Title,
			Descriptions: req.Descriptions,
		}

		if _, err := db.ExecContext(r.Context(), query, notes.ID , notes.Title, notes.Descriptions); err != nil {
			slog.Error("failed inset note", "error", err)
			http.Error(w, "Failed to insert new notes!", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&notes)
	})

	r.Get("/note", func(w http.ResponseWriter, r *http.Request) {
		query := `SELECT id, title, descriptions FROM notes`
		var args []any

		rows, err := db.QueryContext(r.Context(), query, args...)
		if err != nil {
			slog.Error("failed get data notes", "error", err)
			http.Error(w, "failed to get notes", http.StatusInternalServerError)
			return
		}
		defer rows.Close()


		var notes []Notes
		for rows.Next() {
			var note Notes
			rows.Scan(&note.ID, &note.Title, &note.Descriptions)
			notes =  append(notes, note)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(notes)
	})

	srv := &http.Server{
		Addr: ":8080",
		Handler: r,
	}

	srv.ListenAndServe()
}