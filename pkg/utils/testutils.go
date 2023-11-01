package utils

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
)

const projectDirName = "scrapmarketbe"

// LoadEnv loads env vars from .env
func LoadEnv() {
	cwd, _ := os.Getwd()
	rootPath, _ := os.Getwd()
	err := godotenv.Load(string(rootPath) + `/.env`)
	if err != nil {
		log.WithFields(log.Fields{
			"cause": err,
			"cwd":   cwd,
		}).Fatal("Problem loading .env file")

		os.Exit(-1)
	}
}
