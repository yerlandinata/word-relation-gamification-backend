package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/yerlandinata/word-relation-gamification-backend/src/config"
	"github.com/yerlandinata/word-relation-gamification-backend/src/server"
)

func main() {
	if len(os.Args) == 1 {
		server.Serve()
	} else {
		server.RunTask(os.Args[1])
	}
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load environment variables")
	}

	log.SetFlags(log.LstdFlags | log.Llongfile)

	config.InitDB()
	config.InitAppConfig()
	server.RegisterTasks()
}
