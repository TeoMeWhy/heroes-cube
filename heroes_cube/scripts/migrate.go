package main

import (
	"heroes_cube/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	godotenv.Load("../../.env")

	dsn := os.Getenv("MYSQL_DSN")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro ao conectar no banco:", err)
	}

	db.AutoMigrate(
		models.Creature{},
		models.Race{},
		models.Class{},
		models.Item{},
		models.Inventory{})
}
