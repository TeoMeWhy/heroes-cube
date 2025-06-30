package models

import (
	"math"

	"gorm.io/gorm"
)

type Creature struct {
	Id              string     `json:"id" gorm:"primaryKey"`
	Name            string     `json:"name" gorm:"type:varchar(100);not null"`
	RaceName        string     `json:"race_name" gorm:"type:varchar(100);not null;"`
	Race            Race       `json:"race" gorm:"foreignKey:RaceName;references:Name"`
	ClassName       string     `json:"class_name" gorm:"type:varchar(100);not null;"`
	Class           Class      `json:"class" gorm:"foreignKey:ClassName;references:Name"`
	PtsStrength     int        `json:"pts_strength" gorm:"not null;type:int"`
	PtsDexterity    int        `json:"pts_dexterity" gorm:"not null;type:int"`
	PtsIntelligence int        `json:"pts_intelligence" gorm:"not null;type:int"`
	PtsWisdom       int        `json:"pts_wisdom" gorm:"not null;type:int"`
	PtsHitPoints    int        `json:"pts_hit_points" gorm:"not null;type:int"`
	PtsExperience   int        `json:"pts_experience" gorm:"not null;type:int"`
	PtsSkill        int        `json:"pts_skill" gorm:"not null;type:int"`
	InventoryId     string     `json:"inventory_id" gorm:"type:varchar(100);not null;unique"`
	Inventory       *Inventory `json:"inventory" gorm:"foreignKey:InventoryId;references:Id"`
}

func (creature *Creature) GetStrength() int {

	itemsStrength := 0
	for _, item := range creature.Inventory.Items {
		itemsStrength += item.ModStrength
	}
	return creature.PtsStrength + creature.Race.ModStrength + itemsStrength
}

func (creature *Creature) GetDexterity() int {

	itemsDexterity := 0
	for _, item := range creature.Inventory.Items {
		itemsDexterity += item.ModDexterity
	}

	return creature.PtsDexterity + creature.Race.ModDexterity + itemsDexterity
}

func (creature *Creature) GetIntelligence() int {

	itemsIntelligence := 0
	for _, item := range creature.Inventory.Items {
		itemsIntelligence += item.ModIntelligence
	}

	return creature.PtsIntelligence + creature.Race.ModIntelligence + itemsIntelligence
}

func (creature *Creature) GetWisdom() int {

	itemsWisdom := 0
	for _, item := range creature.Inventory.Items {
		itemsWisdom += item.ModWisdom
	}

	return creature.PtsWisdom + creature.Race.ModWisdom + itemsWisdom
}

func (creature *Creature) GetLevel(pts int) int {
	return int(math.Pow(float64(pts/1000), 0.5) + 1.5)
}

func (creature *Creature) CheckLvlUp(pts int) bool {

	lvlNow := creature.GetLevel(creature.PtsExperience)
	lvlNext := creature.GetLevel(creature.PtsExperience + pts)
	return lvlNow < lvlNext
}

func (creature *Creature) SetHitPoints() {
	creature.PtsHitPoints = 10 + creature.GetStrength()
}

func (creature *Creature) GetInventoryDamage() int {

	itemsDamage := 0
	for _, item := range creature.Inventory.Items {
		itemsDamage += item.Damage
	}
	return itemsDamage

}

func (creature *Creature) Save(db *gorm.DB) error {

	if err := db.Save(creature).Error; err != nil {
		return err
	}

	if err := creature.Inventory.Save(db); err != nil {
		return err
	}

	return nil
}

func GetCreature(db *gorm.DB, id string) (*Creature, error) {
	creature := &Creature{Id: id}

	err := db.
		Preload("Race").
		Preload("Class").
		Preload("Inventory").
		Preload("Inventory.Items").
		First(creature).Error

	if err != nil {
		return nil, err
	}

	creature.Inventory.SetSelloutItensPrice()

	return creature, nil
}

func NewCreature(id, name, raceName, className string, db *gorm.DB) (*Creature, error) {

	c, err := GetClass(db, className)
	if err != nil {
		return nil, err
	}

	r, err := GetRace(db, raceName)
	if err != nil {
		return nil, err
	}

	creature := &Creature{
		Id:              id,
		Name:            name,
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

	creature.SetHitPoints()

	return creature, nil
}
