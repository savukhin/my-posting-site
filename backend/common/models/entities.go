package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID        int          `db:"id"`
	Username  string       `db:"username"`
	Email     string       `db:"email"`
	Password  string       `db:"password"`
	AvatarID  int          `db:"avatarId"`
	CreatedAt time.Time    `json:"createdAt"`
	UpdatedAt time.Time    `json:"updatedAt"`
	DeletedAt sql.NullTime `json:"deletedAt"`
}

type FileType string

const (
	TextFile  FileType = "text"
	PhotoFile FileType = "photo"
)

type OwnerType string

const (
	UserOwner OwnerType = "text"
	PostOwner OwnerType = "photo"
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
	AuthorID  string       `db:"authorId"`
	CreatedAt time.Time    `json:"createdAt"`
	UpdatedAt time.Time    `json:"updatedAt"`
	DeletedAt sql.NullTime `json:"deletedAt"`
}
