package domain

import (
	"time"
)

type Domain struct {
	ID        int        `db:"id" json:"id"`
	Label     string     `db:"label" json:"label"`
	CreatedAt *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
}
