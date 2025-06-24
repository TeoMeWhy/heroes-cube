package models

import "gorm.io/gorm"

// classes: Guerreiro, Mago, Clérigo, Ladino
// raças: Humano, Elfo, Anão, Halfling

type Person struct {
	Id              string `json:"id" gorm:"primaryKey"`
	RaceName        string `json:"race_name" gorm:"type:varchar(100);not null;"`
	Race            Race   `json:"race" gorm:"foreignKey:RaceName;references:Name"`
	ClassName       string `json:"class_name" gorm:"type:varchar(100);not null;"`
	Class           Class  `json:"class" gorm:"foreignKey:ClassName;references:Name"`
	PtsStrength     int    `json:"pts_strength" gorm:"not null;type:int"`
	PtsDexterity    int    `json:"pts_dexterity" gorm:"not null;type:int"`
	PtsIntelligence int    `json:"pts_intelligence" gorm:"not null;type:int"`
	PtsWisdom       int    `json:"pts_wisdom" gorm:"not null;type:int"`
}

func (p *Person) GetStrength() int {
	return p.PtsStrength + p.Race.ModStrength
}

func (p *Person) GetDexterity() int {
	return p.PtsDexterity + p.Race.ModDexterity
}

func (p *Person) GetIntelligence() int {
	return p.PtsIntelligence + p.Race.ModIntelligence
}

func (p *Person) GetWisdom() int {
	return p.PtsWisdom + p.Race.ModWisdom
}

func (p *Person) Save(db *gorm.DB) error {
	return db.Save(p).Error
}

func (p *Person) Load(db *gorm.DB) error {
	if err := db.First(p, "id = ?", p.Id).Error; err != nil {
		return err
	}

	r, err := GetRace(db, p.RaceName)
	if err != nil {
		return err
	}
	p.Race = *r

	c, err := GetClass(db, p.ClassName)
	if err != nil {
		return err
	}
	p.Class = *c

	return nil
}

func NewPerson(id, raceName, className string, db *gorm.DB) (*Person, error) {

	c, err := GetClass(db, className)
	if err != nil {
		return nil, err
	}

	r, err := GetRace(db, raceName)
	if err != nil {
		return nil, err
	}

	p := &Person{
		Id:              id,
		RaceName:        raceName,
		ClassName:       className,
		Race:            *r,
		Class:           *c,
		PtsStrength:     c.InitialStrength,
		PtsDexterity:    c.InitialDexterity,
		PtsIntelligence: c.InitialIntelligence,
		PtsWisdom:       c.InitialWisdom,
	}

	return p, nil
}
