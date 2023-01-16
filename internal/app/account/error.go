package account

import (
	"github.com/labstack/echo/v4"
)

var ErrorSpecificExample = echo.HTTPError{Code: 402, Message: "Domain-specific error"}
