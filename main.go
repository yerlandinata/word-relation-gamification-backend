package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/yerlandinata/word-relation-gamification-backend/src/config"
	"github.com/yerlandinata/word-relation-gamification-backend/src/server"
)

func main() {
	server.Serve()
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load environment variables")
	}

	log.SetFlags(log.LstdFlags | log.Llongfile)

	config.InitDB()
	config.InitAppConfig()
}
