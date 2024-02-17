package repository

import (
	_paymentEntity "github.com/Rayato159/isekai-shop-api/modules/payment/entity"
	_paymentException "github.com/Rayato159/isekai-shop-api/modules/payment/exception"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type paymentRepositoryImpl struct {
	db     *gorm.DB
	logger echo.Logger
}

func NewPaymentRepositoryImpl(db *gorm.DB, logger echo.Logger) PaymentRepository {
	return &paymentRepositoryImpl{
		db:     db,
		logger: logger,
	}
}

func (r *paymentRepositoryImpl) InsertPayment(paymentEntity *_paymentEntity.Payment) (*_paymentEntity.Payment, error) {
	insertedPayment := new(_paymentEntity.Payment)

	if err := r.db.Create(paymentEntity).Scan(insertedPayment).Error; err != nil {
		r.logger.Error("Failed to insert payment", err.Error())
		return nil, &_paymentException.InsertPaymentException{}
	}

	return insertedPayment, nil
}

func (r *paymentRepositoryImpl) CalculatePlayerBalance(playerID string) (*_paymentEntity.PlayerBalanceDto, error) {
	balanceDto := new(_paymentEntity.PlayerBalanceDto)

	if err := r.db.Model(
		&_paymentEntity.Payment{},
	).Where(
		"player_id = ?", playerID,
	).Select(
		"player_id, sum(amount) as balance",
	).Group(
		"player_id",
	).Scan(&balanceDto).Error; err != nil {
		r.logger.Error("Failed to calculate player balance", err.Error())
		return nil, &_paymentException.CalculatePlayerBalanceException{PlayerID: playerID}
	}

	return balanceDto, nil
}
