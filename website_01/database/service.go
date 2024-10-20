package database

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	"website-01/internal/models"
)

type Service struct {
	db     *sql.DB
	Models *models.Models
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Initialize() error {
	var err error

	var fsPath = "vault"

	//create a new file at the fsPath called database
	if err = os.MkdirAll(filepath.Join(fsPath, "database"), 0755); err != nil {
		return err
	}

	dbPath := filepath.Join(fsPath, "database", "app.db")
	fmt.Println("Database path: ", dbPath)

	s.db, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}

	if err = sqlite.SetWALMode(s.db); err != nil {
		fmt.Println("Failed to set WAL mode: ", err)
	}

	if err = sqlite.RunMigrations(s.db, "file://migrations"); err != nil {
		return err
	}

	s.Models = models.NewModels(s.db)

	fmt.Println("Database initialized successfully")
	return nil
}

func (s *Service) Close() error {
	if s.db != nil {
		return s.db.Close()
	}
	return nil
}

func (s *Service) GetDB() *sql.DB {
	return s.db
}

func (s *Service) GetModels() *models.Models {
	return s.Models
}
