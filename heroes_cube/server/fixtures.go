package server

import (
	"fmt"
	"heroes_cube/models"
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func mockDB(name string) (*gorm.DB, error) {

	uri := fmt.Sprintf("file:%s?mode=memory", name)
	log.Println(uri)

	db, err := gorm.Open(sqlite.Open(uri), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(
		models.Creature{},
		models.Race{},
		models.Class{},
		models.Item{},
		models.Inventory{})

	data, _ := os.ReadFile("../../data/races.sql")
	races := string(data)
	if err := db.Exec(races).Error; err != nil {
		return nil, err
	}

	data, _ = os.ReadFile("../../data/classes.sql")
	classes := string(data)
	if err := db.Exec(classes).Error; err != nil {
		return nil, err
	}

	data, _ = os.ReadFile("../../data/items.sql")
	items := string(data)
	if err := db.Exec(items).Error; err != nil {
		return nil, err
	}

	creature, err := models.NewCreature("1", "C01", "Humano", "Mago", db)
	if err != nil {
		return nil, err
	}

	espada, err := models.GetItem(db, "1")
	if err != nil {
		return nil, err
	}
	if err := creature.Inventory.AddItem(*espada); err != nil {
		return nil, err
	}

	armadura, err := models.GetItem(db, "3")
	if err != nil {
		return nil, err
	}
	if err := creature.Inventory.AddItem(*armadura); err != nil {
		return nil, err
	}

	botas, err := models.GetItem(db, "13")
	if err != nil {
		return nil, err
	}
	if err := creature.Inventory.AddItem(*botas); err != nil {
		return nil, err
	}

	// cajado, err := models.GetItem(db, "19")
	// if err != nil {
	// 	return nil, err
	// }
	// if err := creature.Inventory.AddItem(*cajado); err != nil {
	// 	return nil, err
	// }

	if err := creature.Save(db); err != nil {
		return nil, err
	}

	return db, nil

}

func mockServer(db *gorm.DB) (*Server, error) {

	controller := &Controller{
		Db: db,
	}

	fiberApp := fiber.New(fiber.Config{
		AppName: "Heroes Cube API",
	})

	validate := validator.New()

	server := &Server{
		Port:       "8080",
		DB:         db,
		App:        fiberApp,
		Controller: controller,
		Validate:   validate,
	}

	return server, nil
}
