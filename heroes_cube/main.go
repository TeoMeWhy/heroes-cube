package main

import (
	"heroes_cube/configs"
	"heroes_cube/server"
	"log"

	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load("../.env")
	conf, err := configs.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	server, err := server.NewServer(conf)
	if err != nil {
		log.Fatal(err)
	}

	server.Start()

}
