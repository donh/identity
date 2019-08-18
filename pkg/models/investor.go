package models

import (
	"time"
)

// Investor represents the investor model
type Investor struct {
	DID     string    `db:"did" validate:"required"`
	Claim   string    `db:"claim"`
	Status  string    `db:"status"`
	Created time.Time `db:"created"`
	Updated time.Time `db:"updated"`
}
