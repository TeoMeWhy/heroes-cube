package models

import "gorm.io/gorm"

type Item struct {
	Id              string `json:"id" gorm:"primaryKey"`
	Name            string `json:"name" gorm:"type:varchar(100);not null;"`
	Description     string `json:"description" gorm:"type:text;"`
	Category        string `json:"category" gorm:"type:varchar(100);not null;"`
	Type            string `json:"type" gorm:"type:varchar(100);not null;"`
	Weight          int    `json:"weight" gorm:"not null;type:int"`
	Price           int    `json:"price" gorm:"not null;type:int"`
	Damage          int    `json:"damage" gorm:"type:int"`
	ModStrength     int    `json:"mod_strength" gorm:"not null;type:int"`
	ModDexterity    int    `json:"mod_dexterity" gorm:"not null;type:int"`
	ModIntelligence int    `json:"mod_intelligence" gorm:"not null;type:int"`
	ModWisdom       int    `json:"mod_wisdom" gorm:"not null;type:int"`
}

func GetItem(db *gorm.DB, id string) (*Item, error) {
	item := &Item{Id: id}
	if err := db.First(item).Error; err != nil {
		return nil, err
	}
	return item, nil
}

func (i *Item) SetSelloutPrice() {
	i.Price = i.Price / 2
}
