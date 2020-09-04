package user

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// TestUserSuccess : hi
func TestUserSuccess(t *testing.T) {
	// Setup
	e := echo.New()
	q := make(url.Values)
	q.Set("name", "123")
	q.Set("tame", "345")
	req := httptest.NewRequest(http.MethodGet, "/?"+q.Encode(), nil) //, strings.NewReader(userJSON))
	// req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, Test(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		// assert.Equal(t, userJSON, rec.Body.String())
	}
}

// TestUserSuccess : hi
func TestUserFail(t *testing.T) {
	// Setup
	e := echo.New()
	q := make(url.Values)
	q.Set("name", "q")
	q.Set("tame", "c")
	req := httptest.NewRequest(http.MethodGet, "/?"+q.Encode(), nil) //, strings.NewReader(userJSON))
	// req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, Test(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}
