package app

import (
	"context"
	"database/sql"
	"errors"

	"github.com/infinitedaremo/go-api-demo/internal/db"
)

var ErrNotFound = errors.New("not found")

type PersonService interface {
	GetPerson(context.Context, int64) (Person, error)
	GetPortfolio(context.Context, int64) (*Portfolio, error)
}

type personServiceImpl struct {
	db *db.Queries
}

func NewPersonService(db *db.Queries) PersonService {
	return &personServiceImpl{db: db}
}

func (p *personServiceImpl) GetPerson(ctx context.Context, id int64) (Person, error) {
	row, err := p.db.GetPerson(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return Person{}, ErrNotFound
		default:
			return Person{}, err
		}
	}

	return Person{
		FirstName: row.FirstName,
		LastName:  row.LastName,
	}, nil
}

func (p *personServiceImpl) GetPortfolio(ctx context.Context, i int64) (*Portfolio, error) {
	portfolio, err := p.db.GetPortfolio(ctx, i)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return PortfolioRowToView(portfolio)
}
