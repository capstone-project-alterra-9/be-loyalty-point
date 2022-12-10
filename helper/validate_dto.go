package helper

import (
	"github.com/labstack/echo/v4"
)

var (
	ErrUsernameLength       = echo.NewHTTPError(400, "Username must be between 8 and 16 characters")
	ErrUsernameAlphanumeric = echo.NewHTTPError(400, "Username must be alphanumeric")
	ErrEmailLength          = echo.NewHTTPError(400, "Email must be at least 8 characters")
	ErrEmailUsername        = echo.NewHTTPError(400, "Email and Username cannot be the same")
	ErrPasswordLength       = echo.NewHTTPError(400, "Password must be at least 8 characters")
	ErrPasswordAlphanumeric = echo.NewHTTPError(400, "Password must be alphanumeric")
	ErrSameDataRequest      = echo.NewHTTPError(400, "Data request cannot be the same")
	ErrEmptyData            = echo.NewHTTPError(400, "Data cannot be empty")
)
