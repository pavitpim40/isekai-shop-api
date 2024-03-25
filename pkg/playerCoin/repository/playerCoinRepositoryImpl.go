package repository

import (
	"github.com/Rayato159/isekai-shop-api/databases"
	"github.com/Rayato159/isekai-shop-api/entities"
	_playerCoin "github.com/Rayato159/isekai-shop-api/pkg/playerCoin/exception"
	_playerCoinModel "github.com/Rayato159/isekai-shop-api/pkg/playerCoin/model"

	"github.com/labstack/echo/v4"
)

type playerCoinImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewPlayerCoinRepositoryImpl(db databases.Database, logger echo.Logger) PlayerCoinRepository {
	return &playerCoinImpl{
		db:     db,
		logger: logger,
	}
}

func (r *playerCoinImpl) Recording(playerCoinEntity *entities.PlayerCoin) (*entities.PlayerCoin, error) {
	insertedPlayerCoin := new(entities.PlayerCoin)

	if err := r.db.ConnectionGetting().Create(playerCoinEntity).Scan(insertedPlayerCoin).Error; err != nil {
		r.logger.Error("Player's balance recording failed:", err.Error())
		return nil, &_playerCoin.CoinAdding{}
	}

	return insertedPlayerCoin, nil
}

func (r *playerCoinImpl) Showing(playerID string) (*_playerCoinModel.PlayerCoinShowing, error) {
	playerCoin := new(_playerCoinModel.PlayerCoinShowing)

	if err := r.db.ConnectionGetting().Model(
		&entities.PlayerCoin{},
	).Where(
		"player_id = ?", playerID,
	).Select(
		"player_id, sum(amount) as coin",
	).Group(
		"player_id",
	).Scan(&playerCoin).Error; err != nil {
		r.logger.Error("Calculating player coin failed:", err.Error())
		return nil, &_playerCoin.PlayerCoinShowing{}
	}

	return playerCoin, nil
}
