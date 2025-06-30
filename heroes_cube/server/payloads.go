package server

import "heroes_cube/models"

type PayloadPostInventoryItem struct {
	ItemID string `json:"item_id" validate:"required"`
}

type PayloadPostCreature struct {
	ID    string `json:"id" validate:"required"`
	Name  string `json:"name" validate:"required"`
	Race  string `json:"race" validate:"required"`
	Class string `json:"class" validate:"required"`
}

type BaseCreature struct {
	Id              string `json:"id"`
	Name            string `json:"name"`
	RaceName        string `json:"race_name"`
	ClassName       string `json:"class_name"`
	PtsStrength     int    `json:"pts_strength"`
	PtsDexterity    int    `json:"pts_dexterity"`
	PtsIntelligence int    `json:"pts_intelligence"`
	PtsWisdom       int    `json:"pts_wisdom"`
	PtsHitPoints    int    `json:"pts_hit_points"`
	PtsExperience   int    `json:"pts_experience"`
	PtsSkill        int    `json:"pts_skill"`
}

func BaseCreatureFromModel(c models.Creature) BaseCreature {

	return BaseCreature{
		Id:              c.Id,
		Name:            c.Name,
		RaceName:        c.RaceName,
		ClassName:       c.ClassName,
		PtsStrength:     c.PtsStrength,
		PtsDexterity:    c.PtsDexterity,
		PtsIntelligence: c.PtsIntelligence,
		PtsWisdom:       c.PtsWisdom,
		PtsHitPoints:    c.PtsHitPoints,
		PtsExperience:   c.PtsExperience,
		PtsSkill:        c.PtsSkill,
	}

}
