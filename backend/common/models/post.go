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
	ID    int      `db:"id"`
	Path  string   `db:"filepath"`
	Title string   `db:"title"`
	Type  FileType `db:"file_type"`
	// Order     int          `db:"order"`
	Owner     OwnerType    `db:"owner_type"`
	OwnerID   int          `db:"owner_id"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}

type Post struct {
	ID        int          `db:"id"`
	AuthorID  int          `db:"author_id"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
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

func GetPostByID(id int) (*Post, []*File, error) {
	post := &Post{}
	query := fmt.Sprintf("SELECT * FROM posts WHERE id = %d LIMIT 1", id)
	if err := DB.Get(post, query); err != nil {
		return nil, nil, err
	}

	files := []*File{}
	query = fmt.Sprintf(`SELECT * FROM files WHERE owner_type='post' AND owner_id=%d`, id)
	err := DB.Select(&files, query)
	if err != nil {
		return nil, nil, err
	}

	return post, files, nil
}
