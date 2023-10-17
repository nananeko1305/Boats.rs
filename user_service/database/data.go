package data

import (
	"gorm.io/gorm"
)

type DataLayer struct {
	dbClient *gorm.DB
}

func InitDataLayer(dbClient *gorm.DB) *DataLayer {
	return &DataLayer{
		dbClient: dbClient,
	}
}
