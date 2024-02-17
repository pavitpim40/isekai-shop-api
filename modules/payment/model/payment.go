package model

import "time"

type (
	Payment struct {
		ID        uint64    `json:"id"`
		PlayerID  string    `json:"playerID"`
		Amount    int64     `json:"amount"`
		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
	}

	TopUpReq struct {
		PlayerID string
		Amount   int64 `json:"amount" validate:"required,gt=0"`
	}

	PlayerBalance struct {
		PlayerID string `json:"playerID"`
		Balance  int64  `json:"balance"`
	}

	BuyItemReq struct {
		PlayerID string
		ItemID   uint64 `json:"itemID" validate:"required,gt=0"`
		Quantity uint   `json:"quantity" validate:"required,gt=0"`
	}

	SellItemReq struct {
		PlayerID string
		ItemID   uint64 `json:"itemID" validate:"required,gt=0"`
		Quantity uint   `json:"quantity" validate:"required,gt=0"`
	}
)
