package backend

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	e := echo.New()
	e.GET("/v1/users/status", HealthCheck())

	req := httptest.NewRequest(http.MethodGet, "/v1/users/status", http.NoBody)
	resp := httptest.NewRecorder()

	e.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("status code was expected %d, but it was obtained %d", http.StatusOK, resp.Code)
	}
}
