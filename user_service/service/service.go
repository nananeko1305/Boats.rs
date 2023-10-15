package service

import (
	data "user_service/database"
	"user_service/models"
)

type BikeService struct {
	dataLayer *data.DataLayer
}

func InitBikeService(dataLayer *data.DataLayer) *BikeService {
	return &BikeService{
		dataLayer: dataLayer,
	}
}

func (service *BikeService) CreateBike(bike models.Bike) {
	service.CreateBike(bike)
}
