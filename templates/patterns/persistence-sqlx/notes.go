package persist

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

// Note is an example entity. Replace / extend with PRO models when assembling.
// Timestamps are strings so SQLite/Postgres/MySQL scan without driver-specific time hooks.
type Note struct {
	ID        int64  `db:"id" json:"id"`
	Title     string `db:"title" json:"title"`
	Body      string `db:"body" json:"body"`
	CreatedAt string `db:"created_at" json:"created_at"`
	UpdatedAt string `db:"updated_at" json:"updated_at"`
}

// Store is a sqlx-backed repository for notes.
type Store struct {
	db *sqlx.DB
}

func NewStore(db *sqlx.DB) *Store {
	return &Store{db: db}
}

func (s *Store) DB() *sqlx.DB { return s.db }

func (s *Store) Create(ctx context.Context, title, body string) (Note, error) {
	switch s.db.DriverName() {
	case "postgres", "sqlite":
		var n Note
		q := `INSERT INTO notes (title, body) VALUES (?, ?) RETURNING id, title, body, created_at, updated_at`
		if err := s.db.GetContext(ctx, &n, s.db.Rebind(q), title, body); err != nil {
			return Note{}, fmt.Errorf("persist: create note: %w", err)
		}
		return n, nil
	default:
		res, err := s.db.ExecContext(ctx, s.db.Rebind(`INSERT INTO notes (title, body) VALUES (?, ?)`), title, body)
		if err != nil {
			return Note{}, fmt.Errorf("persist: create note: %w", err)
		}
		id, err := res.LastInsertId()
		if err != nil {
			return Note{}, fmt.Errorf("persist: create note id: %w", err)
		}
		return s.Get(ctx, id)
	}
}

func (s *Store) Get(ctx context.Context, id int64) (Note, error) {
	var n Note
	err := s.db.GetContext(ctx, &n, s.db.Rebind(`SELECT id, title, body, created_at, updated_at FROM notes WHERE id = ?`), id)
	if err == sql.ErrNoRows {
		return Note{}, fmt.Errorf("persist: note %d: %w", id, err)
	}
	if err != nil {
		return Note{}, err
	}
	return n, nil
}

func (s *Store) List(ctx context.Context, limit int) ([]Note, error) {
	if limit <= 0 {
		limit = 50
	}
	var rows []Note
	err := s.db.SelectContext(ctx, &rows, s.db.Rebind(`SELECT id, title, body, created_at, updated_at FROM notes ORDER BY id DESC LIMIT ?`), limit)
	return rows, err
}

func (s *Store) Update(ctx context.Context, id int64, title, body string) (Note, error) {
	_, err := s.db.ExecContext(ctx, s.db.Rebind(`UPDATE notes SET title = ?, body = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?`), title, body, id)
	if err != nil {
		return Note{}, err
	}
	return s.Get(ctx, id)
}

func (s *Store) Delete(ctx context.Context, id int64) error {
	res, err := s.db.ExecContext(ctx, s.db.Rebind(`DELETE FROM notes WHERE id = ?`), id)
	if err != nil {
		return err
	}
	n, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if n == 0 {
		return sql.ErrNoRows
	}
	return nil
}
