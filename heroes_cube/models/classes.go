package models

import "gorm.io/gorm"

type Class struct {
	Name                string `json:"name" gorm:"type:varchar(100);primaryKey;not null"`
	Description         string `json:"description" gorm:"not null;type:text"`
	InitialStrength     int    `json:"init_strength" gorm:"not null;type int"`
	InitialDexterity    int    `json:"init_dexterity" gorm:"not null;type:int"`
	InitialIntelligence int    `json:"init_intelligence" gorm:"not null;type:int"`
	InitialWisdom       int    `json:"init_wisdom" gorm:"not null;type:int"`
}

func GetClass(db *gorm.DB, name string) (*Class, error) {
	c := &Class{Name: name}
	if err := db.First(c).Error; err != nil {
		return nil, err
	}
	return c, nil
}
