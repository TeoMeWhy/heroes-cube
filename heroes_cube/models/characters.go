package models

type Character interface {
	BaseDamage() int
	GetSpecialAbility() int
}

// Warrior
type Warrior struct {
	*Creature
}

func (w *Warrior) BaseDamage() int {
	return w.GetInventoryDamage() + w.GetStrength()
}

func (w *Warrior) GetSpecialAbility() int {
	return w.BaseDamage() * 2
}

// Rogue
type Rogue struct {
	*Creature
}

func (r *Rogue) BaseDamage() int {
	return r.GetInventoryDamage() + r.GetDexterity()
}

func (r *Rogue) GetSpecialAbility() int {
	return int(float64(r.BaseDamage()) * (1 + (float64(r.GetDexterity()) / 10)))
}

// Mage
type Mage struct {
	*Creature
}

func (m *Mage) BaseDamage() int {
	return m.GetInventoryDamage() + m.GetIntelligence()
}

func (m *Mage) GetSpecialAbility() int {
	return m.BaseDamage() + m.GetIntelligence()
}

// Cleric
type Cleric struct {
	*Creature
}

func (c *Cleric) BaseDamage() int {
	return c.GetInventoryDamage() + c.GetWisdom()
}

func (c *Cleric) GetSpecialAbility() int {
	return c.GetWisdom() * 2
}

// Criando Chars
func NewCharacter(creature *Creature) Character {
	switch creature.Class.Name {
	case "Guerreiro":
		return &Warrior{creature}
	case "Ladino":
		return &Rogue{creature}
	case "Mago":
		return &Mage{creature}
	case "Clérigo":
		return &Cleric{creature}
	default:
		return nil
	}
}
