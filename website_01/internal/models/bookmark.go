package models

import (
	"context"
	"database/sql"
	"time"
)

type Bookmark struct {
	ID         int64     `json:"id"`
	UploadedAt time.Time `json:"uploaded_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Type       string    `json:"type"`
	AssetName  string    `json:"asset_name"`
	Content    string    `json:"content"`
	Source     string    `json:"source"`
	Context    string    `json:"context"`
}

type BookmarkModel struct {
	DB *sql.DB
}

func (m BookmarkModel) Insert(bookmark *Bookmark) error {
	query := `
		INSERT INTO bookmarks (uploaded_at, updated_at, type, asset_name, content, source, context)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`

	args := []interface{}{
		bookmark.UploadedAt,
		bookmark.UpdatedAt,
		bookmark.Type,
		bookmark.AssetName,
		bookmark.Content,
		bookmark.Source,
		bookmark.Context,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := m.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	bookmark.ID = id
	return nil
}

// Add more methods as needed, such as Get, GetAll, Update, Delete, etc.

// GetRecent returns the n most recent bookmarks
func (m BookmarkModel) GetRecent(n int) ([]*Bookmark, error) {
	query := `
        SELECT id, uploaded_at, updated_at, type, asset_name, content, source, context
        FROM bookmarks
        ORDER BY uploaded_at DESC
        LIMIT ?
    `

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query, n)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookmarks []*Bookmark

	for rows.Next() {
		var b Bookmark
		err := rows.Scan(&b.ID, &b.UploadedAt, &b.UpdatedAt, &b.Type, &b.AssetName, &b.Content, &b.Source, &b.Context)
		if err != nil {
			return nil, err
		}
		bookmarks = append(bookmarks, &b)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return bookmarks, nil
}
