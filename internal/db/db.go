package db

import (
	"CalculatorAppBackend/internal/services/calculationService"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=mysecretpassword port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err = db.AutoMigrate(&calculationService.Calculation{}); err != nil {
		return nil, err
	}

	return db, nil
}
