package models

import "gorm.io/gorm"

type Race struct {
	Name            string `json:"name" gorm:"type:varchar(100);primaryKey;not null"`
	Description     string `json:"description" gorm:"not null;type:text"`
	ModStrength     int    `json:"mod_strength" gorm:"not null;type:int"`
	ModDexterity    int    `json:"mod_dexterity" gorm:"not null;type:int"`
	ModIntelligence int    `json:"mod_intelligence" gorm:"not null;type:int"`
	ModWisdom       int    `json:"mod_wisdom" gorm:"not null;type:int"`
}

func GetRace(db *gorm.DB, name string) (*Race, error) {
	r := &Race{Name: name}
	if err := db.First(r).Error; err != nil {
		return nil, err
	}
	return r, nil
}
