package models

import (
	"heroes_cube/clients/db"
	"heroes_cube/configs"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCharsBaseDamage(t *testing.T) {
	conf := configs.Config{
		DbUser: "admin",
		DbPass: "admin",
		DbHost: "127.0.0.1",
		DbPort: "3306",
		DbName: "heroes",
	}

	db, err := db.GetMySqlClient(conf)
	assert.NoError(t, err)

	creatureWarrior, err := NewCreature("Warrior", "Warrior-humano", "humano", "guerreiro", db)
	assert.NoError(t, err)
	charWarrior := NewCharacter(creatureWarrior)

	creatureRogue, err := NewCreature("Rogue", "Rogue-humano", "humano", "ladino", db)
	assert.NoError(t, err)
	charRogue := NewCharacter(creatureRogue)

	creatureMage, err := NewCreature("Mage", "Mage-humano", "humano", "mago", db)
	assert.NoError(t, err)
	charMage := NewCharacter(creatureMage)

	creatureCleric, err := NewCreature("Cleric", "Cleric-humano", "humano", "clérigo", db)
	assert.NoError(t, err)
	charCleric := NewCharacter(creatureCleric)

	pairTests := []struct {
		name           string
		char           Character
		expectedDamage int
		expectError    error
	}{
		{
			name:           "Warrior Base Damage",
			char:           charWarrior,
			expectedDamage: 6,
			expectError:    nil,
		},
		{
			name:           "Rogue Base Damage",
			char:           charRogue,
			expectedDamage: 7,
			expectError:    nil,
		},
		{
			name:           "Mage Base Damage",
			char:           charMage,
			expectedDamage: 7,
			expectError:    nil,
		},
		{
			name:           "Cleric Base Damage",
			char:           charCleric,
			expectedDamage: 7,
			expectError:    nil,
		},
	}

	for _, tt := range pairTests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expectedDamage, tt.char.BaseDamage())
			assert.Equal(t, tt.expectError, err)
		})
	}

}

func TestCharsBaseDamageWithItems(t *testing.T) {
	conf := configs.Config{
		DbUser: "admin",
		DbPass: "admin",
		DbHost: "127.0.0.1",
		DbPort: "3306",
		DbName: "heroes",
	}

	db, err := db.GetMySqlClient(conf)
	assert.NoError(t, err)

	espadaLonga, err := GetItem(db, "1")
	assert.NoError(t, err)

	botasAgilidade, err := GetItem(db, "5")
	assert.NoError(t, err)

	cajadoFogo, err := GetItem(db, "7")
	assert.NoError(t, err)

	cajadoVida, err := GetItem(db, "14")
	assert.NoError(t, err)

	creatureWarrior, err := NewCreature("Warrior", "Warrior-humano", "humano", "guerreiro", db)
	assert.NoError(t, err)
	err = creatureWarrior.Inventory.AddItem(*espadaLonga)
	assert.NoError(t, err)
	charWarrior := NewCharacter(creatureWarrior)

	creatureRogue, err := NewCreature("Rogue", "Rogue-humano", "humano", "ladino", db)
	assert.NoError(t, err)
	err = creatureRogue.Inventory.AddItem(*botasAgilidade)
	assert.NoError(t, err)
	charRogue := NewCharacter(creatureRogue)

	creatureMage, err := NewCreature("Mage", "Mage-humano", "humano", "mago", db)
	assert.NoError(t, err)
	err = creatureMage.Inventory.AddItem(*cajadoFogo)
	assert.NoError(t, err)
	charMage := NewCharacter(creatureMage)

	creatureCleric, err := NewCreature("Cleric", "Cleric-humano", "humano", "clérigo", db)
	assert.NoError(t, err)
	err = creatureCleric.Inventory.AddItem(*cajadoVida)
	assert.NoError(t, err)
	charCleric := NewCharacter(creatureCleric)

	pairTests := []struct {
		name           string
		char           Character
		expectedDamage int
		expectError    error
	}{
		{
			name:           "Warrior Base Damage",
			char:           charWarrior,
			expectedDamage: 16,
			expectError:    nil,
		},
		{
			name:           "Rogue Base Damage",
			char:           charRogue,
			expectedDamage: 9,
			expectError:    nil,
		},
		{
			name:           "Mage Base Damage",
			char:           charMage,
			expectedDamage: 19,
			expectError:    nil,
		},
		{
			name:           "Cleric Base Damage",
			char:           charCleric,
			expectedDamage: 16,
			expectError:    nil,
		},
	}

	for _, tt := range pairTests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expectedDamage, tt.char.BaseDamage())
			assert.Equal(t, tt.expectError, err)
		})
	}

}

func TestCharsSpecialAbility(t *testing.T) {
	conf := configs.Config{
		DbUser: "admin",
		DbPass: "admin",
		DbHost: "127.0.0.1",
		DbPort: "3306",
		DbName: "heroes",
	}

	db, err := db.GetMySqlClient(conf)
	assert.NoError(t, err)

	espadaLonga, err := GetItem(db, "1")
	assert.NoError(t, err)

	botasAgilidade, err := GetItem(db, "5")
	assert.NoError(t, err)

	cajadoFogo, err := GetItem(db, "7")
	assert.NoError(t, err)

	cajadoVida, err := GetItem(db, "14")
	assert.NoError(t, err)

	creatureWarrior, err := NewCreature("Warrior", "Warrior-humano", "humano", "guerreiro", db)
	assert.NoError(t, err)
	err = creatureWarrior.Inventory.AddItem(*espadaLonga)
	assert.NoError(t, err)
	charWarrior := NewCharacter(creatureWarrior)

	creatureRogue, err := NewCreature("Rogue", "Rogue-humano", "humano", "ladino", db)
	assert.NoError(t, err)
	err = creatureRogue.Inventory.AddItem(*botasAgilidade)
	assert.NoError(t, err)
	charRogue := NewCharacter(creatureRogue)

	creatureMage, err := NewCreature("Mage", "Mage-humano", "humano", "mago", db)
	assert.NoError(t, err)
	err = creatureMage.Inventory.AddItem(*cajadoFogo)
	assert.NoError(t, err)
	charMage := NewCharacter(creatureMage)

	creatureCleric, err := NewCreature("Cleric", "Cleric-humano", "humano", "clérigo", db)
	assert.NoError(t, err)
	err = creatureCleric.Inventory.AddItem(*cajadoVida)
	assert.NoError(t, err)
	charCleric := NewCharacter(creatureCleric)

	pairTests := []struct {
		name           string
		char           Character
		expectedDamage int
		expectError    error
	}{
		{
			name:           "Warrior Base Damage",
			char:           charWarrior,
			expectedDamage: 32,
			expectError:    nil,
		},
		{
			name:           "Rogue Base Damage",
			char:           charRogue,
			expectedDamage: 17,
			expectError:    nil,
		},
		{
			name:           "Mage Base Damage",
			char:           charMage,
			expectedDamage: 28,
			expectError:    nil,
		},
		{
			name:           "Cleric Base Damage",
			char:           charCleric,
			expectedDamage: 18,
			expectError:    nil,
		},
	}

	for _, tt := range pairTests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expectedDamage, tt.char.GetSpecialAbility())
			assert.Equal(t, tt.expectError, err)
		})
	}

}
