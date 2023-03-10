package models

import "github.com/uptrace/bun"

type User struct {
	bun.BaseModel

	ID       []rune
	Email    string
	Password string
}
