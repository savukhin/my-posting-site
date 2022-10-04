package models

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"gopkg.in/validator.v2"
)

type User struct {
	ID        int32         `db:"id"`
	Username  string        `db:"username" validate:"regexp=^[a-zA-Z]{1}([\\w@\\s.]+){5}$"`
	Email     string        `db:"email" validate:"regexp=^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\\.[a-zA-Z0-9-.]+$"`
	Password  string        `db:"password" validate:"min=8,regexp=^.*[a-zA-Z].*$,regexp=^.*\\d.*$"`
	AvatarID  sql.NullInt32 `db:"avatar_id"`
	CreatedAt time.Time     `db:"created_at"`
	UpdatedAt time.Time     `db:"updated_at"`
	DeletedAt sql.NullTime  `db:"deleted_at"`
}

func CreateUser(username, email, password string) (*User, error) {
	user := &User{
		Username: username,
		Email:    email,
		Password: password,
	}

	return user, user.Valid()
}

func (user *User) ComparePassword(unhashedPassword string) error {
	if user.Password != unhashedPassword {
		return errors.New("passwords not equal")
	}
	return nil
}

func (user *User) Valid() error {
	if err := validator.Validate(user); err != nil {
		return err
	}
	fmt.Println("password ok")
	return nil
}

func (user *User) update() (*User, error) {
	if err := user.Valid(); err != nil {
		return nil, err
	}

	createUser := `UPDATE users SET username=?, email=?, password=? WHERE id=?`
	DB.MustExec(createUser, user.Username, user.Email, user.Password, user.ID)
	return GetUserByUsername(user.Username)
}

func (user *User) create() (*User, error) {
	if err := user.Valid(); err != nil {
		return nil, err
	}

	sqlStr := fmt.Sprintf(`insert into users(username, email, password) values ('%s', '%s', '%s')`,
		user.Username, user.Email, user.Password)

	if _, err := DB.Exec(sqlStr); err != nil {
		return nil, err
	}

	return GetUserByUsername(user.Username)
}

func (user *User) Save() (*User, error) {
	if _, err := GetUserByUsername(user.Username); err == nil {
		return user.update()
	}

	return user.create()
}

func (user *User) IsNotExists() error {
	query := fmt.Sprintf("SELECT * FROM users WHERE username='%s' OR email='%s' LIMIT 1", user.Username, user.Email)
	result := &User{}
	err := DB.Get(&result, query)
	if result == nil || err == nil {
		return nil
	}
	if result.Username == user.Username {
		return errors.New("user with this username is already exists")
	}
	return errors.New("user with this email is already exists")
}

func GetUserByUsername(username string) (*User, error) {
	user := &User{}
	query := fmt.Sprintf("SELECT * FROM users WHERE username='%s' LIMIT 1", username)
	err := DB.Get(user, query)

	return user, err
}
