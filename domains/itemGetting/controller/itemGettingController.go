package controller

import "github.com/labstack/echo/v4"

type ItemGettingController interface {
	Listing(pctx echo.Context) error
}
