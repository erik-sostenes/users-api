package user

import (
	"bytes"
	"fmt"
	"github.com/erik-sostenes/users-api/internal/mooc/user/business/domain"
	"github.com/erik-sostenes/users-api/internal/mooc/user/business/services/create"
	"github.com/erik-sostenes/users-api/internal/mooc/user/infrastructure/persistence"
	"github.com/erik-sostenes/users-api/internal/shared/domain/bus/command"
	"github.com/labstack/echo/v4"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

type funcHandler func() (Handler, error)

func TestUser_Create(t *testing.T) {
	tsc := map[string]struct {
		request            *http.Request
		userHandler        funcHandler
		expectedStatusCode int
	}{
		"given an existing valid .csv file, a status code 201 is expected": {
			request: func() *http.Request {
				data, err := os.ReadFile("users.csv")
				if err != nil {
					t.Fatal(err)
				}

				var body bytes.Buffer
				writer := multipart.NewWriter(&body)
				part, _ := writer.CreateFormFile("users", "users.csv")
				part.Write(data)
				writer.Close()

				req := httptest.NewRequest(http.MethodPost, "/v1/users", &body)
				req.Header.Set("Content-Type", writer.FormDataContentType())
				req.Header.Set("Content-Length", fmt.Sprint(body.Len()))

				return req
			}(),
			userHandler: func() (handler Handler, err error) {
				commandHandler := &create.CreateUserCommandHandler{
					create.UserCreator{
						persistence.NewMockUserRepository[domain.UserId, domain.User](),
					},
				}

				bus := make(command.CommandBus[create.UserCommand])
				if err = bus.Record(create.UserCommand{}, commandHandler); err != nil {
					return
				}

				return NewUserHandler(&bus), nil
			},
			expectedStatusCode: http.StatusCreated,
		},
		"given a valid .csv file not existing, a status code 404 is expected": {
			request: func() *http.Request {
				data, err := os.ReadFile("users.csv")
				if err != nil {
					t.Fatal(err)
				}

				var body bytes.Buffer
				writer := multipart.NewWriter(&body)
				part, _ := writer.CreateFormFile("some_file", "users.csv")
				part.Write(data)
				writer.Close()

				req := httptest.NewRequest(http.MethodPost, "/v1/users", &body)
				req.Header.Set("Content-Type", writer.FormDataContentType())
				req.Header.Set("Content-Length", fmt.Sprint(body.Len()))

				return req
			}(),
			userHandler: func() (handler Handler, err error) {
				commandHandler := &create.CreateUserCommandHandler{}

				bus := make(command.CommandBus[create.UserCommand])
				if err = bus.Record(create.UserCommand{}, commandHandler); err != nil {
					return
				}

				return NewUserHandler(&bus), nil
			},
			expectedStatusCode: http.StatusBadRequest,
		},
	}

	for name, ts := range tsc {
		t.Run(name, func(t *testing.T) {
			userHandler, err := ts.userHandler()
			if err != nil {
				t.Skip(err)
			}

			engine := echo.New()
			engine.POST("/v1/users", userHandler.Create())

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
