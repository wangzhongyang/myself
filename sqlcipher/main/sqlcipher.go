package main

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"

	_ "github.com/xeodou/go-sqlcipher"
)

type SQLCipher struct {
	*sql.DB
}

func (s *SQLCipher) OpenFile(fileName, param string) error {
	err := errors.New("")
	s.DB, err = sql.Open("sqlite3", fmt.Sprintf("%s?%s", fileName, param))
	return err
}

func (s *SQLCipher) Exec(query string) (sql.Result, error) {
	return s.DB.Exec(query)
}

type SQLCipherOrm struct {
	*gorm.DB
}

type User struct {
	gorm.Model
	ID        int
	Name      string
	Password  string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt time.Time
}

func (s *SQLCipherOrm) OpenFile(fileName, param string) error {
	var err error
	s.DB, err = gorm.Open("sqlite3", fmt.Sprintf("%s?%s", fileName, param))
	return err
}
