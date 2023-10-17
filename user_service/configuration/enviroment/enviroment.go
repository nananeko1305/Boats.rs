package configuration

import "os"

type Enviroment struct {
	SERVER_PORT string
	DB_HOST     string
	DB_PASS     string
	DB_USER     string
	DB_PORT     string
	DB_NAME     string
}

func NewEnviroment() *Enviroment {
	return &Enviroment{
		SERVER_PORT: os.Getenv("SERVER_PORT"),
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_PASS:     os.Getenv("DB_PASS"),
		DB_USER:     os.Getenv("DB_USER"),
		DB_PORT:     os.Getenv("DB_PORT"),
		DB_NAME:     os.Getenv("DB_NAME"),
	}
}
