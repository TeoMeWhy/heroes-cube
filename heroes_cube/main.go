package main

import (
	"fmt"
	"heroes_cube/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	godotenv.Load("../.env")

	dsn := os.Getenv("MYSQL_DSN")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro ao conectar no banco:", err)
	}

	db.AutoMigrate(
		models.Person{},
		models.Race{},
		models.Class{},
		models.Item{},
		models.Inventory{})

	// person, err := models.NewPerson("fodase", "human", "warrior", db)
	// if err != nil {
	// 	log.Fatal("Erro ao criar pessoa:", err)
	// }
	// log.Println("Pessoa criada:", person)

	// item01, _ := models.GetItem(db, "1")
	// item03, _ := models.GetItem(db, "3")

	// person.Inventory = &models.Inventory{Id: person.Id, Items: []models.Item{*item01, *item03}}

	// person.Save(db)

	// if err := person.Save(db); err != nil {
	// 	log.Fatal("Erro ao salvar pessoa:", err)
	// }
	// log.Println("Pessoa salva no banco:", person)

	p, _ := models.GetPerson(db, "fodase")
	log.Println("Pessoa carregada do banco:", p)

	log.Println("Inventário carregado do banco:", *p.Inventory)

	fmt.Println("Heroes Cube")
}
