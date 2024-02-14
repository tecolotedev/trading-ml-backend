package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type EnvVarsType struct {
	PORT        string
	DB_HOST     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	IS_LOCAL    bool

	EMAIL_HOST     string
	EMAIL_PORT     string
	EMAIL_USER     string
	EMAIL_PASSWORD string

	FRONT_URL string
}

var EnvVars EnvVarsType

// Initiate env vars
func SetUpConfig() {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Error loading .env file: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	EnvVars.PORT = port

	//connection to postgres database
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	isLocal, err := strconv.ParseBool(os.Getenv("IS_LOCAL"))

	if err != nil {
		isLocal = true
	}

	EnvVars.DB_HOST = dbHost
	EnvVars.DB_USER = dbUser
	EnvVars.DB_PASSWORD = dbPassword
	EnvVars.DB_NAME = dbName
	EnvVars.IS_LOCAL = isLocal

	//connection to mailtrap
	emailHost := os.Getenv("EMAIL_HOST")
	emailPassword := os.Getenv("EMAIL_PASSWORD")
	emailUser := os.Getenv("EMAIL_USER")
	emailPort := os.Getenv("EMAIL_PORT")

	EnvVars.EMAIL_HOST = emailHost
	EnvVars.EMAIL_PASSWORD = emailPassword
	EnvVars.EMAIL_USER = emailUser
	EnvVars.EMAIL_PORT = emailPort

	//front url
	frontUrl := os.Getenv("FRONT_URL")
	EnvVars.FRONT_URL = frontUrl
}
