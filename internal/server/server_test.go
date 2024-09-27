package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

var _ ServerInterface = (*Server)(nil)

func Test_Ping(t *testing.T) {
	s, err := NewServer(zap.NewNop())
	assert.NoError(t, err)

	req, err := http.NewRequest("GET", "/ping", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
