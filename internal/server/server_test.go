package server

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/infinitedaremo/go-api-demo/internal/app"
	mocks "github.com/infinitedaremo/go-api-demo/internal/app/_mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
)

var _ ServerInterface = (*Server)(nil)

func Test_Ping(t *testing.T) {
	s, err := NewServer(zap.NewNop(), nil)
	assert.NoError(t, err)

	req, err := http.NewRequest("GET", "/ping", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestServer_GetPortfolio(t *testing.T) {
	ps := mocks.NewPersonService(t)
	ps.EXPECT().GetPortfolio(mock.Anything, int64(1)).Return(&app.Portfolio{
		Person: app.Person{
			FirstName: "gavin woods",
			LastName:  "test",
		},
		WorkExperience: []app.WorkExperience{
			{CompanyName: "test", JobTitle: "test guy", Tenure: "2020 - 2024"},
		},
	}, nil)

	// Test generic errors
	ps.EXPECT().GetPortfolio(mock.Anything, int64(2)).Return(nil, errors.New("not found"))

	s, err := NewServer(zap.NewNop(), ps)
	assert.NoError(t, err)

	req, err := http.NewRequest("GET", "/portfolio/1", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	req, err = http.NewRequest("GET", "/portfolio/2", nil)
	assert.NoError(t, err)

	w = httptest.NewRecorder()
	s.router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
