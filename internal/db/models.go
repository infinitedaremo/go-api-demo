// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"database/sql"
	"time"
)

type Person struct {
	ID        int64
	FirstName string
	LastName  string
}

type WorkExperience struct {
	ID          int64
	PersonID    int64
	CompanyName string
	JobTitle    string
	StartDate   time.Time
	EndDate     sql.NullTime
}
