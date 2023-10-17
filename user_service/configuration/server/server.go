package server

import (
	connection "user_service/configuration/database"
	configuration "user_service/configuration/enviroment"
	data "user_service/database"
	"user_service/handlers"
	"user_service/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Start(enviroment *configuration.Enviroment) {

	engine := gin.Default()
	engine.ForwardedByClientIP = true
	engine.SetTrustedProxies([]string{"127.0.0.1"})

	// This method create Three Layer Architecture and connect each layer
	initLayers(engine, enviroment)

	// Listen on SERVER_PORT
	engine.Run(enviroment.SERVER_PORT)
}

func initLayers(engine *gin.Engine, enviroment *configuration.Enviroment) {

	// Init Connection to Postgres
	collection := connection.ConnectToDatabase(enviroment)

	// Define logger
	logger := zap.Logger{}

	// Create all layers and connect it
	dataLayer := data.InitDataLayer(collection)
	serviceLayer := service.InitBikeService(dataLayer)
	handlerLayer := handlers.InitHandler(serviceLayer, &logger)

	// Init routes
	handlers.SetupRoutes(engine, handlerLayer)

}
