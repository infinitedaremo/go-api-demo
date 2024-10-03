package app

import (
	"database/sql"
	"testing"
	"time"

	"github.com/infinitedaremo/go-api-demo/internal/db"
	"github.com/stretchr/testify/assert"
)

func TestPortfolioRowToView(t *testing.T) {
	// Mock data for portfolio
	portfolioRows := []db.GetPortfolioRow{
		{
			Person: db.Person{
				FirstName: "John",
				LastName:  "Doe",
			},
			WorkExperience: db.WorkExperience{
				CompanyName: "Company A",
				JobTitle:    "Software Engineer",
				StartDate:   time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
				EndDate:     sql.NullTime{Valid: true, Time: time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC)},
			},
		},
		{
			Person: db.Person{
				FirstName: "John",
				LastName:  "Doe",
			},
			WorkExperience: db.WorkExperience{
				CompanyName: "Company B",
				JobTitle:    "Senior Engineer",
				StartDate:   time.Date(2021, time.February, 1, 0, 0, 0, 0, time.UTC),
				EndDate:     sql.NullTime{Valid: false}, // Current job
			},
		},
	}

	// Call the function
	portfolio, err := PortfolioRowToView(portfolioRows)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, portfolio)

	// Check person details
	assert.Equal(t, "John", portfolio.Person.FirstName)
	assert.Equal(t, "Doe", portfolio.Person.LastName)

	// Check work experience details
	assert.Len(t, portfolio.WorkExperience, 2)
	assert.Equal(t, "Company A", portfolio.WorkExperience[0].CompanyName)
	assert.Equal(t, "Software Engineer", portfolio.WorkExperience[0].JobTitle)
	assert.Equal(t, "Jan 2019 - Jan 2021", portfolio.WorkExperience[0].Tenure)

	assert.Equal(t, "Company B", portfolio.WorkExperience[1].CompanyName)
	assert.Equal(t, "Senior Engineer", portfolio.WorkExperience[1].JobTitle)
	assert.Equal(t, "Feb 2021 - Present", portfolio.WorkExperience[1].Tenure)

}
