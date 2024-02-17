package main

import (
	"github.com/Rayato159/isekai-shop-api/config"
	"github.com/Rayato159/isekai-shop-api/databases"
	_adminEntity "github.com/Rayato159/isekai-shop-api/modules/admin/entity"
	_inventoryEntity "github.com/Rayato159/isekai-shop-api/modules/inventory/entity"
	_itemEntity "github.com/Rayato159/isekai-shop-api/modules/item/entity"
	_orderEntity "github.com/Rayato159/isekai-shop-api/modules/order/entity"
	_paymentEntity "github.com/Rayato159/isekai-shop-api/modules/payment/entity"
	_playerEntity "github.com/Rayato159/isekai-shop-api/modules/player/entity"
)

func main() {
	appConfig := config.GetAppConfig()
	database := databases.NewPostgresDatabase(appConfig.DatabaseConfig)

	uuidMigreate(database)
	playerMigrate(database)
	adminMigrate(database)
	itemMigrate(database)
	paymentMigrate(database)
	inventoryMigrate(database)
	orderMigrate(database)
}

func uuidMigreate(db databases.Database) {
	db.GetDb().Raw(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`).Scan(&struct{}{})
}

func playerMigrate(db databases.Database) {
	db.GetDb().Migrator().CreateTable(&_playerEntity.Player{})
}

func adminMigrate(db databases.Database) {
	db.GetDb().Migrator().CreateTable(&_adminEntity.Admin{})
}

func itemMigrate(db databases.Database) {
	db.GetDb().Migrator().CreateTable(&_itemEntity.Item{})
}

func paymentMigrate(db databases.Database) {
	db.GetDb().Migrator().CreateTable(&_paymentEntity.Payment{})
}

func inventoryMigrate(db databases.Database) {
	db.GetDb().Migrator().CreateTable(&_inventoryEntity.Inventory{})
}

func orderMigrate(db databases.Database) {
	db.GetDb().Migrator().CreateTable(&_orderEntity.Order{})
}
