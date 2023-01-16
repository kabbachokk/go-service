package errors

import "github.com/labstack/echo"

var BadInputError = echo.HTTPError{Code: 400, Message: "Bad input"}
var InternalServerError = echo.HTTPError{Code: 500, Message: "Server error"}
