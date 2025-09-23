package main

import (
	"heroes_cube/clients/db"
	"heroes_cube/configs"
	"heroes_cube/models"
	"log"

	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load("../../.env")

	conf, err := configs.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	con, err := db.GetMySqlClient(*conf)
	if err != nil {
		log.Fatal(err)
	}

	con.AutoMigrate(
		models.Creature{},
		models.Race{},
		models.Class{},
		models.Item{},
		models.Inventory{})
}
