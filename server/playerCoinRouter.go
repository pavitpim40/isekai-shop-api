package server

import (
	_playerCoinController "github.com/Rayato159/isekai-shop-api/pkg/playerCoin/controller"
	_playerCoinRepository "github.com/Rayato159/isekai-shop-api/pkg/playerCoin/repository"
	_playerCoinService "github.com/Rayato159/isekai-shop-api/pkg/playerCoin/service"
	"github.com/Rayato159/isekai-shop-api/server/customMiddleware"
)

func (s *echoServer) initPlayerCoinRouter(customMiddleware customMiddleware.CustomMiddleware) {
	router := s.baseRouter.Group("/player-coin")

	playerCoinRepository := _playerCoinRepository.NewPlayerCoinRepositoryImpl(s.db, s.app.Logger)

	playerCoinService := _playerCoinService.NewPlayerCoinServiceImpl(
		playerCoinRepository,
	)
	playerCoinController := _playerCoinController.NewPlayerCoinControllerImpl(playerCoinService)

	router.POST("", playerCoinController.CoinAdding, customMiddleware.PlayerAuthorizing)
	router.GET("", playerCoinController.PlayerCoinShowing, customMiddleware.PlayerAuthorizing)
}
