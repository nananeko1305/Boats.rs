package handlers

import (
	"user_service/models"

	"github.com/gin-gonic/gin"
)

func (handler *BikeHandler) CreateBike(context *gin.Context) {

	var requestBody models.Bike
	if err := context.ShouldBindJSON(&requestBody); err != nil {
		// handler.logger.Warn("error in parsing json",
		// 	zapcore.Field{Key: "error", String: err.Error()})

		// Ovojiti!!!
		// context.Status(http.StatusInternalServerError)
		context.Header("Content-Type", "application/json")
		// Create a response JSON object
		response := gin.H{
			"message": "error in parsing json",
			"data":    err.Error(),
		}

		// Return the response as JSON
		context.JSON(500, response)
		return
	}

	context.JSON(200, "OK")

}
