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

func (controller *Controller) GetCreatures() ([]BaseCreature, error) {

	creatures := []BaseCreature{}
	modelCreatures := []models.Creature{}
	if err := controller.Db.Find(&modelCreatures).Error; err != nil {
		return nil, err
	}

	for _, c := range modelCreatures {
		creatures = append(creatures, BaseCreatureFromModel(c))
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

func (controller *Controller) DeleteCreature(id string) error {

	creature, err := models.GetCreature(controller.Db, id)
	if err != nil {
		return err
	}

	return controller.Db.Delete(creature).Error
}

func (controller *Controller) PostInventoryItem(inventory_id, item_id string) error {

	inventory, err := models.GetInventory(controller.Db, inventory_id)
	if err != nil {
		return err
	}

	item, err := models.GetItem(controller.Db, item_id)
	if err != nil {
		return err
	}

	if err := inventory.AddItem(*item); err != nil {
		return err
	}

	inventory.Save(controller.Db)
	return nil

}
