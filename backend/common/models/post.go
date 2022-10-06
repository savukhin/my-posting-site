package models

import (
	"database/sql"
	"fmt"
	"time"
)

type FileType string

const (
	TextFile  FileType = "text"
	PhotoFile FileType = "photo"
)

type OwnerType string

const (
	UserOwner OwnerType = "user"
	PostOwner OwnerType = "post"
)

type File struct {
	ID        int          `db:"id"`
	Path      string       `db:"path"`
	Title     string       `db:"title"`
	Type      FileType     `db:"fileType"`
	Order     int          `db:"order"`
	Owner     OwnerType    `db:"ownerType"`
	OwnerID   int          `db:"ownerId"`
	CreatedAt time.Time    `json:"createdAt"`
	UpdatedAt time.Time    `json:"updatedAt"`
	DeletedAt sql.NullTime `json:"deletedAt"`
}

type Post struct {
	ID        int          `db:"id"`
	AuthorID  int          `db:"authorId"`
	CreatedAt time.Time    `json:"createdAt"`
	UpdatedAt time.Time    `json:"updatedAt"`
	DeletedAt sql.NullTime `json:"deletedAt"`
}

func (file *File) Save() error {
	sqlStr := fmt.Sprintf(
		`INSERT INTO files (filepath, title, file_type, owner_type, owner_id)
		VALUES ('%s', '%s', '%s', '%s', %d) RETURNING id`,
		file.Path, file.Title, file.Type, file.Owner, file.OwnerID,
	)

	var fileId int
	err := DB.QueryRow(sqlStr).Scan(&fileId)

	return err
}

func (post *Post) Save(files []*File) error {
	sqlStr := fmt.Sprintf(
		`INSERT INTO posts (author_id) VALUES (%d) RETURNING id`,
		post.AuthorID,
	)

	var postId int
	err := DB.QueryRow(sqlStr).Scan(&postId)
	if err != nil {
		return err
	}

	post.ID = postId

	for _, file := range files {
		file.OwnerID = postId
		file.Save()
	}

	return nil
}
