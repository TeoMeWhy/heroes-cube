package models

import (
	"heroes_cube/clients/db"
	"heroes_cube/configs"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInventoryAddItem(t *testing.T) {

	conf := configs.Config{
		DbUser: "admin",
		DbPass: "admin",
		DbHost: "127.0.0.1",
		DbPort: "3306",
		DbName: "heroes",
	}

	db, err := db.GetMySqlClient(conf)
	assert.NoError(t, err)

	item01, _ := GetItem(db, "1")
	item02, _ := GetItem(db, "2")
	item03, _ := GetItem(db, "3")
	item04, _ := GetItem(db, "4")
	item05, _ := GetItem(db, "5")
	item06, _ := GetItem(db, "6")

	pairTests := []struct {
		name        string
		inventory   Inventory
		items       []Item
		expectError error
	}{
		{
			name:        "Add item 1 sucessfullyto inventory",
			inventory:   Inventory{Id: "inv01"},
			items:       []Item{*item01},
			expectError: nil,
		},
		{
			name:        "Add 2 items sucessfully to inventory",
			inventory:   Inventory{Id: "inv02"},
			items:       []Item{*item01, *item03},
			expectError: nil,
		},
		{
			name:        "Add 3 items sucessfully to inventory",
			inventory:   Inventory{Id: "inv03"},
			items:       []Item{*item01, *item03, *item05},
			expectError: nil,
		},
		{
			name:        "Add 2 items not sucessfully to inventory",
			inventory:   Inventory{Id: "inv04", Items: []Item{*item01}},
			items:       []Item{*item02},
			expectError: ErrorItemTypeAlreadyExists,
		},
		{
			name:        "Add 6 items not sucessfully to inventory",
			inventory:   Inventory{Id: "inv04", Items: []Item{*item01, *item03, *item05}},
			items:       []Item{*item02, *item04, *item06},
			expectError: ErrorItemTypeAlreadyExists,
		},
	}

	for _, tt := range pairTests {
		t.Run(tt.name, func(t *testing.T) {
			for _, item := range tt.items {
				err := tt.inventory.AddItem(item)
				assert.ErrorIs(t, err, tt.expectError)
			}
		})
	}

}

func TestInventoryRemoveItem(t *testing.T) {

	conf := configs.Config{
		DbUser: "admin",
		DbPass: "admin",
		DbHost: "127.0.0.1",
		DbPort: "3306",
		DbName: "heroes",
	}

	db, err := db.GetMySqlClient(conf)
	assert.NoError(t, err)

	item01, _ := GetItem(db, "1")
	item02, _ := GetItem(db, "2")

	pairTests := []struct {
		name        string
		inventory   Inventory
		items       []Item
		expectError error
	}{
		{
			name:        "Remove item 1 sucessfully from inventory",
			inventory:   Inventory{Id: "inv01", Items: []Item{*item01, *item02}},
			items:       []Item{*item01},
			expectError: nil,
		},
		{
			name:        "Remove item 1 sucessfully from inventory",
			inventory:   Inventory{Id: "inv01", Items: []Item{*item02}},
			items:       []Item{*item01},
			expectError: ErrorItemIdNotFoundOnInventory,
		},
	}

	for _, tt := range pairTests {
		t.Run(tt.name, func(t *testing.T) {
			for _, item := range tt.items {
				err := tt.inventory.RemoveItem(item.Id)
				assert.ErrorIs(t, err, tt.expectError)
			}
		})
	}

}
