package config

import (
	"fmt"

	"github.com/flipmorsch/Blincast/modules/channel"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(models ...interface{}) (*gorm.DB, error) {
	dbPath := ":memory:"

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("open sqlite db: %w", err)
	}

	if len(models) > 0 {
		if err := db.AutoMigrate(models...); err != nil {
			return nil, fmt.Errorf("automigrate models: %w", err)
		}
	}

	DB = db
	db.AutoMigrate(channel.ChannelEntity{})
	return DB, nil
}

func CloseDB() error {
	if DB == nil {
		return nil
	}
	sqlDB, err := DB.DB()
	if err != nil {
		return fmt.Errorf("get sql DB from gorm: %w", err)
	}
	return sqlDB.Close()
}
