package models

import (
	"time"
)

// Email represents the email model
type Email struct {
	Email   string    `db:"email" validate:"required"`
	Code    string    `db:"code"`
	Valid   int64     `db:"valid"`
	Created time.Time `db:"created"`
	Updated time.Time `db:"updated"`
}
