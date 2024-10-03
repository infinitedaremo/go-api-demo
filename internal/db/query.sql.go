// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createPerson = `-- name: CreatePerson :one
INSERT INTO person (first_name, last_name)
VALUES (?, ?)
RETURNING id, first_name, last_name
`

type CreatePersonParams struct {
	FirstName string
	LastName  string
}

func (q *Queries) CreatePerson(ctx context.Context, arg CreatePersonParams) (Person, error) {
	row := q.db.QueryRowContext(ctx, createPerson, arg.FirstName, arg.LastName)
	var i Person
	err := row.Scan(&i.ID, &i.FirstName, &i.LastName)
	return i, err
}

const createWorkExperience = `-- name: CreateWorkExperience :exec
INSERT INTO work_experience (person_id, company_name, job_title, start_date, end_date)
VALUES (?, ?, ?, ?, ?)
`

type CreateWorkExperienceParams struct {
	PersonID    int64
	CompanyName string
	JobTitle    string
	StartDate   time.Time
	EndDate     sql.NullTime
}

func (q *Queries) CreateWorkExperience(ctx context.Context, arg CreateWorkExperienceParams) error {
	_, err := q.db.ExecContext(ctx, createWorkExperience,
		arg.PersonID,
		arg.CompanyName,
		arg.JobTitle,
		arg.StartDate,
		arg.EndDate,
	)
	return err
}

const getPerson = `-- name: GetPerson :one
SELECT first_name, last_name FROM person
WHERE id = ? LIMIT 1
`

type GetPersonRow struct {
	FirstName string
	LastName  string
}

func (q *Queries) GetPerson(ctx context.Context, id int64) (GetPersonRow, error) {
	row := q.db.QueryRowContext(ctx, getPerson, id)
	var i GetPersonRow
	err := row.Scan(&i.FirstName, &i.LastName)
	return i, err
}

const getPortfolio = `-- name: GetPortfolio :many
SELECT person.id, person.first_name, person.last_name, work_experience.id, work_experience.person_id, work_experience.company_name, work_experience.job_title, work_experience.start_date, work_experience.end_date
FROM person
JOIN work_experience ON work_experience.person_id = person.id
WHERE work_experience.person_id = ?
ORDER BY work_experience.start_date DESC
LIMIT 25
`

type GetPortfolioRow struct {
	Person         Person
	WorkExperience WorkExperience
}

func (q *Queries) GetPortfolio(ctx context.Context, personID int64) ([]GetPortfolioRow, error) {
	rows, err := q.db.QueryContext(ctx, getPortfolio, personID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetPortfolioRow
	for rows.Next() {
		var i GetPortfolioRow
		if err := rows.Scan(
			&i.Person.ID,
			&i.Person.FirstName,
			&i.Person.LastName,
			&i.WorkExperience.ID,
			&i.WorkExperience.PersonID,
			&i.WorkExperience.CompanyName,
			&i.WorkExperience.JobTitle,
			&i.WorkExperience.StartDate,
			&i.WorkExperience.EndDate,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
