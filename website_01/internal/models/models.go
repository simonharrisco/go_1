package models

import (
	"database/sql"
)

type Models struct {
	Bookmarks BookmarkModel
}

func NewModels(db *sql.DB) *Models {
	return &Models{
		Bookmarks: BookmarkModel{DB: db},
	}
}
