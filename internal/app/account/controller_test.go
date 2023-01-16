package account

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	e          = echo.New()
	uc         = UseCaseInterface(new(useCaseMock))
	controller = ControllerInterface(NewController(&uc))
	authCtx    = &AuthCtx{ID: 123}
)

func init() {
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(AuthCtxKey, authCtx)
			return next(c)
		}
	})

	Handle(e, &controller)
}

func TestCreate(t *testing.T) {
	rb := &requestBody{FirstName: "John", LastName: "Doe", BirthDate: "2000-05-08"}
	obj, err := rb.convert()
	assert.NoError(t, err)

	(uc.(*useCaseMock)).On("CreateAccount", obj).Return(
		&Response{Id: fmt.Sprint(authCtx.ID)},
		nil,
	)

	body, err := json.Marshal(rb)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/account", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	e.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestGet(t *testing.T) {
	(uc.(*useCaseMock)).On("GetAccount", authCtx.ID).Return(
		&Response{Id: fmt.Sprint(authCtx.ID)},
		nil,
	)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/account/123", nil)
	req.Header.Set("Content-Type", "application/json")
	e.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestUpdate(t *testing.T) {
	rb := &requestBody{FirstName: "John", LastName: "Doe", BirthDate: "2000-05-08"}
	obj, err := rb.convert()
	assert.NoError(t, err)

	(uc.(*useCaseMock)).On("UpdateAccount", authCtx.ID, obj).Return(
		&Response{Id: fmt.Sprint(authCtx.ID)},
		nil,
	)

	body, err := json.Marshal(rb)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/account/123", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	e.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestDelete(t *testing.T) {
	(uc.(*useCaseMock)).On("DeleteAccount", authCtx.ID).Return(nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/account/123", nil)
	req.Header.Set("Content-Type", "application/json")
	e.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
