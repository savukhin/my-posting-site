package models

import (
	"os"
	"path/filepath"
)

func RunSchema(path string) error {
	c, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	query := string(c)

	stmt, err := DB.Prepare(query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec()

	return err
}

func CreateTables() error {
	folder := filepath.Join("backend", "common", "sql")
	files, err := os.ReadDir(folder)
	if err != nil {
		return err
	}
	for _, f := range files {
		path := filepath.Join(folder, f.Name())
		err := RunSchema(path)
		if err != nil {
			continue
		}
	}
	return nil
}
