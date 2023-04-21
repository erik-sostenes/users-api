package user

import (
	"bytes"
	"fmt"
	"github.com/labstack/echo/v4"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestUser_Create(t *testing.T) {
	tsc := map[string]struct {
		request            *http.Request
		user               Handler
		expectedStatusCode int
	}{
		"given an existing valid .csv file, a status code 201 is expected": {
			user: NewUserHandler(),
			request: func() *http.Request {
				f, err := os.ReadFile("users.csv")
				if err != nil {
					t.Fatal(err)
				}

				var body bytes.Buffer
				writer := multipart.NewWriter(&body)
				part, _ := writer.CreateFormFile("users", "users.csv")
				part.Write(f)
				writer.Close()

				req := httptest.NewRequest(http.MethodPost, "/v1/users", &body)
				req.Header.Set("Content-Type", writer.FormDataContentType())
				req.Header.Set("Content-Length", fmt.Sprint(body.Len()))

				return req
			}(),
			expectedStatusCode: http.StatusCreated,
		},
		"given a valid .csv file not existing, a status code 404 is expected": {
			request: func() *http.Request {
				f, err := os.ReadFile("users.csv")
				if err != nil {
					t.Fatal(err)
				}

				var body bytes.Buffer
				writer := multipart.NewWriter(&body)
				part, _ := writer.CreateFormFile("some_file", "users.csv")
				part.Write(f)
				writer.Close()

				req := httptest.NewRequest(http.MethodPost, "/v1/users", &body)
				req.Header.Set("Content-Type", writer.FormDataContentType())
				req.Header.Set("Content-Length", fmt.Sprint(body.Len()))

				return req
			}(),
			user:               NewUserHandler(),
			expectedStatusCode: http.StatusBadRequest,
		},
	}

	for name, ts := range tsc {
		t.Run(name, func(t *testing.T) {
			engine := echo.New()

			engine.POST("/v1/users", ts.user.Create())

			req := ts.request
			resp := httptest.NewRecorder()
			engine.ServeHTTP(resp, req)

			if resp.Code != ts.expectedStatusCode {
				t.Log(resp.Body.String())
				t.Errorf("status code was expected %d, but it was obtained %d", http.StatusCreated, resp.Code)
			}
		})
	}
}
