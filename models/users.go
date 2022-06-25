package models

import "github.com/uptrace/bun"

type User struct {
	bun.BaseModel `bun:"table:users"`
	ID            int64  `bun:",pk,autoincrement"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	Username      string `json:"username"`
	Password      string `json:"password"`
}
