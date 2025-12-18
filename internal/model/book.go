package model

import "time"

type Book struct {
	ID int `db:"id"`
	Title string `db:"title"`
	Author string `db:"author"`
	CreatedAt time.Time `db:"created_at"`
}