package server

import (
	"heroes_cube/models"

	"gorm.io/gorm"
)

type Controller struct {
	Db *gorm.DB
}

func (controller *Controller) GetRaces() ([]models.Race, error) {
	races := []models.Race{}
	if err := controller.Db.Find(&races).Error; err != nil {
		return nil, err
	}
	return races, nil
}

func (controller *Controller) GetClasses() ([]models.Class, error) {
	classes := []models.Class{}
	if err := controller.Db.Find(&classes).Error; err != nil {
		return nil, err
	}
	return classes, nil
}

func (controller *Controller) GetItems() ([]models.Item, error) {
	items := []models.Item{}
	if err := controller.Db.Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (controller *Controller) GetCreatures() ([]models.Creature, error) {
	creatures := []models.Creature{}
	if err := controller.Db.Find(&creatures).Error; err != nil {
		return nil, err
	}
	return creatures, nil
}

func (controller *Controller) GetCreaturesByID(id string) (*models.Creature, error) {
	creature, err := models.GetCreature(controller.Db, id)
	if err != nil {
		return nil, err
	}
	return creature, nil
}

func (controller *Controller) PostCreature(payload PayloadPostCreature) (*models.Creature, error) {

	_, err := models.GetCreature(controller.Db, payload.ID)
	if err == nil {
		return nil, models.ErrorCreatureAlreadyExists
	}

	creature, err := models.NewCreature(payload.ID, payload.Name, payload.Race, payload.Class, controller.Db)
	if err != nil {
		return nil, err
	}

	if err := creature.Save(controller.Db); err != nil {
		return nil, err
	}
	return creature, nil
}
