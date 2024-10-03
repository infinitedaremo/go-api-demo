package app

import (
	"errors"
	"fmt"

	"github.com/infinitedaremo/go-api-demo/internal/db"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type WorkExperience struct {
	CompanyName string `json:"company_name"`
	JobTitle    string `json:"job_title"`
	Tenure      string `json:"tenure"`
}

type Portfolio struct {
	Person         Person           `json:"person"`
	WorkExperience []WorkExperience `json:"experience"`
}

func PortfolioRowToView(portfolio []db.GetPortfolioRow) (*Portfolio, error) {
	if len(portfolio) == 0 {
		return nil, errors.New("invalid portfolio")
	}

	pf := Portfolio{
		Person: Person{
			FirstName: portfolio[0].Person.FirstName,
			LastName:  portfolio[0].Person.LastName,
		},
		WorkExperience: make([]WorkExperience, len(portfolio)),
	}

	for i, row := range portfolio {
		startDate := row.WorkExperience.StartDate.Format("Jan 2006")
		endDate := "Present"
		if row.WorkExperience.EndDate.Valid {
			endDate = row.WorkExperience.EndDate.Time.Format("Jan 2006")
		}

		pf.WorkExperience[i] = WorkExperience{
			CompanyName: row.WorkExperience.CompanyName,
			JobTitle:    row.WorkExperience.JobTitle,
			Tenure:      fmt.Sprintf("%s - %s", startDate, endDate),
		}
	}

	return &pf, nil
}
