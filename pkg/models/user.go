package models

import (
	"database/sql"
	"time"
)

// User represents the user model
type User struct {
	FirstName string         `db:"firstname" validate:"required"`
	LastName  string         `db:"lastname" validate:"required"`
	Email     string         `db:"email" validate:"required"`
	Phone     string         `db:"phone" validate:"required"`
	Birthday  string         `db:"birthday" validate:"required"`
	SSN       sql.NullString `db:"ssn"`
	Country   sql.NullString `db:"country"`
	Region    sql.NullString `db:"region"`
	City      sql.NullString `db:"city"`
	Street    sql.NullString `db:"street"`
	Zip       int64          `db:"zip"`
	Passport  string         `db:"passport"`
	Address   sql.NullString `db:"address"`
	DID       sql.NullString `db:"did"`
	Created   time.Time      `db:"created"`
	Updated   time.Time      `db:"updated"`
}
