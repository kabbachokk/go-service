package account

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.service/internal/pkg/errors"
)

type Date struct {
	string
}

func (d *Date) Time() (date time.Time, err error) {
	date, err = time.Parse("2006-01-02", d.string)
	return
}

type requestBody struct {
	FirstName string `json:"first_name" form:"first_name"`
	LastName  string `json:"last_name" form:"last_name"`
	BirthDate string `json:"birth_date" form:"birth_date"`
}

func (r *requestBody) convert() (req *Request, err error) {
	date := &Date{r.BirthDate}
	t, err := date.Time()
	if err != nil {
		return
	}
	req = &Request{
		FirstName:  r.FirstName,
		LastName:   r.LastName,
		BirthYear:  t.Year(),
		BirthMonth: int(t.Month()),
		BirthDay:   t.Day(),
	}
	return
}

type ControllerInterface interface {
	CreateAccount(c echo.Context) error
	GetAccount(c echo.Context) error
	UpdateAccount(c echo.Context) error
	DeleteAccount(c echo.Context) error
}

type Controller struct {
	uc *UseCaseInterface
}

func NewController(uc *UseCaseInterface) *Controller {
	return &Controller{
		uc,
	}
}

var _ ControllerInterface = (*Controller)(nil) //check

// will be moved to auth pkg

type AuthCtx struct {
	ID int
}

var AuthCtxKey string = "Auth"

// ---

func (c *Controller) CreateAccount(ctx echo.Context) error {
	rb := new(requestBody)
	if err := ctx.Bind(&rb); err != nil {
		return &errors.BadInputError
	}

	obj, err := rb.convert()
	if err != nil {
		return &errors.InternalServerError
	}

	account, err := (*c.uc).CreateAccount(obj)
	if err != nil {
		return &errors.InternalServerError
	}

	return ctx.JSON(http.StatusOK, newAccountView(account))
}

func (c *Controller) GetAccount(ctx echo.Context) error {
	authCtx := ctx.Get(AuthCtxKey).(*AuthCtx)

	account, err := (*c.uc).GetAccount(authCtx.ID)
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}

	return ctx.JSON(http.StatusOK, newAccountView(account))
}

func (c *Controller) UpdateAccount(ctx echo.Context) error {
	rb := new(requestBody)
	if err := ctx.Bind(&rb); err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}

	obj, err := rb.convert()
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}

	authCtx := ctx.Get(AuthCtxKey).(*AuthCtx)

	account, err := (*c.uc).UpdateAccount(authCtx.ID, obj)
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}

	return ctx.JSON(http.StatusOK, newAccountView(account))
}

func (c *Controller) DeleteAccount(ctx echo.Context) error {
	authCtx := ctx.Get(AuthCtxKey).(*AuthCtx)

	if err := (*c.uc).DeleteAccount(authCtx.ID); err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}

	return ctx.NoContent(http.StatusOK)
}
