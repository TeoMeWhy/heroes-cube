package models

import "gorm.io/gorm"

// classes: Guerreiro, Mago, Clérigo, Ladino
// raças: Humano, Elfo, Anão, Halfling

type Person struct {
	Id              string     `json:"id" gorm:"primaryKey"`
	RaceName        string     `json:"race_name" gorm:"type:varchar(100);not null;"`
	Race            Race       `json:"race" gorm:"foreignKey:RaceName;references:Name"`
	ClassName       string     `json:"class_name" gorm:"type:varchar(100);not null;"`
	Class           Class      `json:"class" gorm:"foreignKey:ClassName;references:Name"`
	PtsStrength     int        `json:"pts_strength" gorm:"not null;type:int"`
	PtsDexterity    int        `json:"pts_dexterity" gorm:"not null;type:int"`
	PtsIntelligence int        `json:"pts_intelligence" gorm:"not null;type:int"`
	PtsWisdom       int        `json:"pts_wisdom" gorm:"not null;type:int"`
	PtsHitPoints    int        `json:"pts_hit_points" gorm:"not null;type:int"`
	InventoryId     string     `json:"inventory_id" gorm:"type:varchar(100);not null;unique"`
	Inventory       *Inventory `json:"inventory" gorm:"foreignKey:InventoryId;references:Id"`
}

func (p *Person) GetStrength() int {

	itemsStrength := 0
	for _, item := range p.Inventory.Items {
		itemsStrength += item.ModStrength
	}
	return p.PtsStrength + p.Race.ModStrength + itemsStrength
}

func (p *Person) GetDexterity() int {

	itemsDexterity := 0
	for _, item := range p.Inventory.Items {
		itemsDexterity += item.ModDexterity
	}

	return p.PtsDexterity + p.Race.ModDexterity + itemsDexterity
}

func (p *Person) GetIntelligence() int {

	itemsIntelligence := 0
	for _, item := range p.Inventory.Items {
		itemsIntelligence += item.ModIntelligence
	}

	return p.PtsIntelligence + p.Race.ModIntelligence + itemsIntelligence
}

func (p *Person) GetWisdom() int {

	itemsWisdom := 0
	for _, item := range p.Inventory.Items {
		itemsWisdom += item.ModWisdom
	}

	return p.PtsWisdom + p.Race.ModWisdom + itemsWisdom
}

func (p *Person) SetHitPoints() {
	p.PtsHitPoints = 10 + p.GetStrength()
}

func (p *Person) GetDamage() int {

	itemsDamage := 0
	for _, item := range p.Inventory.Items {
		itemsDamage += item.Damage
	}
	return itemsDamage

}

func (p *Person) Save(db *gorm.DB) error {

	if err := db.Save(p).Error; err != nil {
		return err
	}

	if err := p.Inventory.Save(db); err != nil {
		return err
	}

	return nil
}

func GetPerson(db *gorm.DB, id string) (*Person, error) {
	p := &Person{Id: id}

	err := db.
		Preload("Race").
		Preload("Class").
		Preload("Inventory").
		Preload("Inventory.Items").
		First(p).Error

	if err != nil {
		return nil, err
	}

	return p, nil
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
		Inventory:       NewInventory(id),
	}

	p.SetHitPoints()

	return p, nil
}
