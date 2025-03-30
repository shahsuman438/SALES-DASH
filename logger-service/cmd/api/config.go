package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
)

var Cnfg Configs

type Configs struct {
	DBUser            string
	DBPassword        string
	DBHost            string
	DBName            string
	ServerPort        string
	ServerHost        string
	ServerEnvironment string
	ServerMode        string
	AppURL            string
}

func init() {
	loadConfig()
}

func loadConfig() {
	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(workingDir)
	serverRootDir := strings.Replace(workingDir, "cmd", "", 1)
	err = godotenv.Load(filepath.Join(serverRootDir, ".env"))

	if err != nil {
		log.Fatal("Error loading .env file, using system environment variables")
	}

	Cnfg = Configs{
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBHost:     os.Getenv("DB_HOST"),
		DBName:     os.Getenv("DB_NAME"),
	}
}
