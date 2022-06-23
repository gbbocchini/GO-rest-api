package models

import "github.com/uptrace/bun"

type Book struct {
	bun.BaseModel `bun:"table:books"`
	ID            int64  `bun:",pk,autoincrement"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	Year          int64  `json:"year"`
}
