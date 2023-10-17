package main

import (
	"log"
	configuration "user_service/configuration/enviroment"
	"user_service/configuration/server"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	enviroment := configuration.NewEnviroment()
	log.Println(enviroment)
	server.Start(enviroment)
}
