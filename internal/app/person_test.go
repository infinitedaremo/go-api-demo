package app

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/infinitedaremo/go-api-demo/internal/db"
	"github.com/stretchr/testify/assert"
)

func Test_GetPerson(t *testing.T) {
	ctx := context.Background()
	dbx, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	// Setup mocked person
	columns := []string{"first_name", "last_name"}
	mock.ExpectQuery("SELECT first_name, last_name FROM person").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows(columns).
			FromCSVString(fmt.Sprintf("%s,%s", "gavin", "woods")))

	mock.ExpectQuery("SELECT first_name, last_name FROM person").
		WithArgs(2).
		WillReturnError(sql.ErrNoRows)

	svc := NewPersonService(db.New(dbx))
	p, err := svc.GetPerson(ctx, 1)
	assert.NoError(t, err)
	assert.Equal(t, "gavin", p.FirstName)
	assert.Equal(t, "woods", p.LastName)

	// Check no records
	p, err = svc.GetPerson(ctx, 2)
	assert.ErrorIs(t, err, ErrNotFound)
}

func Test_personServiceImpl_GetPortfolio(t *testing.T) {
	ctx := context.Background()
	dbx, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	expectedQuery := `SELECT person.id, person.first_name, person.last_name, work_experience.id, work_experience.person_id, work_experience.company_name, work_experience.job_title, work_experience.start_date, work_experience.end_date
FROM person
JOIN work_experience ON work_experience.person_id = person.id
WHERE work_experience.person_id`
	columns := []string{"person.id", "person.first_name", "person.last_name", "work_experience.id", "work_experience.person_id", "work_experience.company_name", "work_experience.job_title", "work_experience.start_date", "work_experience.end_date"}
	mock.ExpectQuery(expectedQuery).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows(columns).
			AddRow(1, "Gavin", "Woods", 1, 1, "Test Company", "Engineer",
				time.Date(2022, 12, 24, 0, 0, 0, 0, time.UTC),
				time.Date(2023, 12, 24, 0, 0, 0, 0, time.UTC),
			).AddRow(1, "Gavin", "Woods", 2, 1, "Test Company 2", "Senior Engineer",
			time.Date(2023, 12, 24, 0, 0, 0, 0, time.UTC),
			nil))

	mock.ExpectQuery(expectedQuery).
		WithArgs(2).
		WillReturnError(sql.ErrNoRows)

	svc := NewPersonService(db.New(dbx))
	p, err := svc.GetPortfolio(ctx, 1)
	assert.NoError(t, err)
	assert.Equal(t, "Gavin", p.Person.FirstName)
	assert.Equal(t, "Dec 2022 - Dec 2023", p.WorkExperience[0].Tenure)
	assert.Equal(t, "Dec 2023 - Present", p.WorkExperience[1].Tenure)

	// Check no records
	p, err = svc.GetPortfolio(ctx, 2)
	assert.ErrorIs(t, err, ErrNotFound)

}
