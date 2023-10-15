package configuration

import "os"

type Enviroment struct {
	SERVER_PORT  string
	USER_DB_HOST string
	USER_DB_PORT string
	DATABASE     string
	COLLECTION   string
}

func NewEnviroment() *Enviroment {
	return &Enviroment{
		SERVER_PORT:  os.Getenv("SERVER_PORT"),
		USER_DB_HOST: os.Getenv("USER_DB_HOST"),
		USER_DB_PORT: os.Getenv("USER_DB_PORT"),
		DATABASE:     os.Getenv("DATABASE"),
		COLLECTION:   os.Getenv("COLLECTION"),
	}
}
