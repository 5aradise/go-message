package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	gormDB *gorm.DB
}

func New(dsn string) (*Database, error) {
	gormDB, err := gorm.Open(sqlite.Open(dsn))
	if err != nil {
		return nil, err
	}

	err = gormDB.AutoMigrate(&User{})
	if err != nil {
		return nil, err
	}

	return &Database{gormDB}, nil
}
