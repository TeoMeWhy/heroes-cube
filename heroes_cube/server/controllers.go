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
