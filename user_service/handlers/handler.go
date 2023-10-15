package handlers

import (
	"user_service/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BikeHandler struct {
	service *service.BikeService
	logger  *zap.Logger
}

func InitHandler(service *service.BikeService, logger *zap.Logger) *BikeHandler {
	return &BikeHandler{
		service: service,
		logger:  logger,
	}
}

func SetupRoutes(ginRouter *gin.Engine, handler *BikeHandler) {

	// Init routes for bikes
	bikeRouter := ginRouter.Group("/bikes")
	bikeRouter.GET("/:id", handler.GetBike)
	bikeRouter.POST("/", handler.CreateBike)
	// bikeRouter.PUT("/:id", handler.UpdateBike)
	// bikeRouter.DELETE("/:id", handler.DeleteBike)

}
