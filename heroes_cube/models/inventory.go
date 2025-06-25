package models

import (
	"slices"

	"gorm.io/gorm"
)

type Inventory struct {
	Id    string `json:"id" gorm:"primaryKey;type:varchar(100)"`
	Items []Item `json:"items" gorm:"many2many:inventory_items;"`
}

func (i *Inventory) GetWeight() int {
	totalWeight := 0
	for _, item := range i.Items {
		totalWeight += item.Weight
	}
	return totalWeight
}

func (i *Inventory) AddItem(item Item) error {
	for _, existingItem := range i.Items {
		if existingItem.Type == item.Type {
			return ErrorItemTypeAlreadyExists
		}
	}

	i.Items = append(i.Items, item)
	return nil
}

func (i *Inventory) RemoveItem(id string) error {

	for ix, v := range i.Items {
		if v.Id == id {
			i.Items = slices.Delete(i.Items, ix, ix+1)
			return nil
		}
	}

	return ErrorItemIdNotFoundOnInventory
}

func (i *Inventory) Save(db *gorm.DB) error {
	if err := db.Save(i).Error; err != nil {
		return err
	}

	if err := db.Model(i).Association("Items").Replace(i.Items); err != nil {
		return err
	}

	return nil
}

func GetInventory(db *gorm.DB, id string) (*Inventory, error) {
	inv := &Inventory{Id: id}
	if err := db.Preload("Items").First(inv).Error; err != nil {
		return nil, err
	}
	return inv, nil
}

func NewInventory(id string) *Inventory {
	return &Inventory{Id: id}
}
