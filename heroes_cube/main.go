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

	db.AutoMigrate(models.Person{}, models.Race{}, models.Class{})

	// person, err := models.NewPerson("fodase", "human", "warrior", db)
	// if err != nil {
	// 	log.Fatal("Erro ao criar pessoa:", err)
	// }
	// log.Println("Pessoa criada:", person)

	// if err := person.Save(db); err != nil {
	// 	log.Fatal("Erro ao salvar pessoa:", err)
	// }
	// log.Println("Pessoa salva no banco:", person)

	p := &models.Person{
		Id: "fodase",
	}

	p.Load(db)
	log.Println("Pessoa carregada do banco:", p)

	fmt.Println("Heroes Cube")
}
