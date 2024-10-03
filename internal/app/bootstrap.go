package app

import (
	"context"
	"database/sql"
	"time"

	"github.com/infinitedaremo/go-api-demo/internal/db"
)

// Bootstrap is used for pre-populating in memory DB with persistent cache
func Bootstrap(ctx context.Context, d *db.Queries) error {
	person, err := d.CreatePerson(ctx, db.CreatePersonParams{
		FirstName: "Gavin",
		LastName:  "Woods",
	})
	if err != nil {
		return err
	}

	exp := []db.CreateWorkExperienceParams{{
		PersonID:    person.ID,
		CompanyName: "Blockdaemon",
		JobTitle:    "Senior Engineering Manager",
		StartDate:   time.Date(2023, 01, 01, 0, 0, 0, 0, time.UTC),
	}, {
		PersonID:    person.ID,
		CompanyName: "Blockdaemon",
		JobTitle:    "Engineering Manager",
		StartDate:   time.Date(2022, 01, 01, 0, 0, 0, 0, time.UTC),
		EndDate: sql.NullTime{
			Time:  time.Date(2023, 01, 01, 0, 0, 0, 0, time.UTC),
			Valid: true,
		},
	}, {
		PersonID:    person.ID,
		CompanyName: "Blockdaemon",
		JobTitle:    "Lead Developer",
		StartDate:   time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC),
		EndDate: sql.NullTime{
			Time:  time.Date(2022, 01, 01, 0, 0, 0, 0, time.UTC),
			Valid: true,
		},
	}, {
		PersonID:    person.ID,
		CompanyName: "Blockdaemon",
		JobTitle:    "Full Stack Developer",
		StartDate:   time.Date(2019, 01, 01, 0, 0, 0, 0, time.UTC),
		EndDate: sql.NullTime{
			Time:  time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC),
			Valid: true,
		},
	}, {
		PersonID:    person.ID,
		CompanyName: "HSBC",
		JobTitle:    "Senior Mulesoft Consultant",
		StartDate:   time.Date(2018, 01, 01, 0, 0, 0, 0, time.UTC),
		EndDate: sql.NullTime{
			Time:  time.Date(2019, 01, 01, 0, 0, 0, 0, time.UTC),
			Valid: true,
		},
	}, {
		PersonID:    person.ID,
		CompanyName: "Atom Bank",
		JobTitle:    "Lead Middleware Developer",
		StartDate:   time.Date(2018, 01, 01, 0, 0, 0, 0, time.UTC),
		EndDate: sql.NullTime{
			Time:  time.Date(2018, 01, 01, 0, 0, 0, 0, time.UTC),
			Valid: true,
		},
	}, {
		PersonID:    person.ID,
		CompanyName: "Atom Bank",
		JobTitle:    "Senior Mulesoft Developer",
		StartDate:   time.Date(2016, 01, 01, 0, 0, 0, 0, time.UTC),
		EndDate: sql.NullTime{
			Time:  time.Date(2018, 01, 01, 0, 0, 0, 0, time.UTC),
			Valid: true,
		},
	}, {
		PersonID:    person.ID,
		CompanyName: "Atom Bank",
		JobTitle:    "Senior Applications Developer",
		StartDate:   time.Date(2015, 01, 01, 0, 0, 0, 0, time.UTC),
		EndDate: sql.NullTime{
			Time:  time.Date(2016, 01, 01, 0, 0, 0, 0, time.UTC),
			Valid: true,
		},
	}, {
		PersonID:    person.ID,
		CompanyName: "Laughing Jackal",
		JobTitle:    "Games Developer",
		StartDate:   time.Date(2011, 01, 01, 0, 0, 0, 0, time.UTC),
		EndDate: sql.NullTime{
			Time:  time.Date(2015, 01, 01, 0, 0, 0, 0, time.UTC),
			Valid: true,
		},
	}}
	for _, xp := range exp {
		err = d.CreateWorkExperience(ctx, xp)
		if err != nil {
			return err
		}
	}

	return nil
}
