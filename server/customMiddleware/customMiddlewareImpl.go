package customMiddleware

import (
	_oauth2Controller "github.com/Rayato159/isekai-shop-api/pkg/oauth2/controller"

	"github.com/Rayato159/isekai-shop-api/config"
	"github.com/labstack/echo/v4"
)

type customMiddlewaresImpl struct {
	oauth2Conf       *config.OAuth2Config
	logger           echo.Logger
	oauth2Controller _oauth2Controller.OAuth2Controller
}

func NewCustomMiddlewaresImpl(
	oauth2Controller _oauth2Controller.OAuth2Controller,
	oauth2Conf *config.OAuth2Config,
	logger echo.Logger,
) CustomMiddleware {
	return &customMiddlewaresImpl{
		oauth2Controller: oauth2Controller,
		oauth2Conf:       oauth2Conf,
		logger:           logger,
	}
}

func (m *customMiddlewaresImpl) PlayerAuthorizing(next echo.HandlerFunc) echo.HandlerFunc {
	return func(pctx echo.Context) error {
		return m.oauth2Controller.PlayerAuthorizing(pctx, next)
	}
}

func (m *customMiddlewaresImpl) AdminAuthorizing(next echo.HandlerFunc) echo.HandlerFunc {
	return func(pctx echo.Context) error {
		return m.oauth2Controller.AdminAuthorizing(pctx, next)
	}
}
