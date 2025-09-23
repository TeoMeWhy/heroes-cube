package models

import (
	"heroes_cube/clients/db"
	"heroes_cube/configs"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImportRaces(t *testing.T) {

	conf := configs.Config{
		DbUser: "admin",
		DbPass: "admin",
		DbHost: "127.0.0.1",
		DbPort: "3306",
		DbName: "heroes",
	}

	db, err := db.GetMySqlClient(conf)
	assert.NoError(t, err)

	cases := []struct {
		name                    string
		raceName                string
		expectedRaceName        string
		expectedModStrength     int
		expectedModDexterity    int
		expectedModIntelligence int
		expectedModWisdom       int
		expectError             error
	}{
		{
			name:                    "Importar raça Humano",
			raceName:                "Humano",
			expectedRaceName:        "Humano",
			expectedModStrength:     1,
			expectedModDexterity:    1,
			expectedModIntelligence: 1,
			expectedModWisdom:       1,
			expectError:             nil,
		},
		{
			name:                    "Importar raça Elfo",
			raceName:                "Elfo",
			expectedRaceName:        "Elfo",
			expectedModStrength:     0,
			expectedModDexterity:    2,
			expectedModIntelligence: 1,
			expectedModWisdom:       1,
			expectError:             nil,
		},
		{
			name:                    "Importar raça Anão",
			raceName:                "Anão",
			expectedRaceName:        "Anão",
			expectedModStrength:     2,
			expectedModDexterity:    1,
			expectedModIntelligence: 0,
			expectedModWisdom:       1,
			expectError:             nil,
		},
		{
			name:                    "Importar raça Hobbit",
			raceName:                "Hobbit",
			expectedRaceName:        "Hobbit",
			expectedModStrength:     0,
			expectedModDexterity:    2,
			expectedModIntelligence: 0,
			expectedModWisdom:       2,
			expectError:             nil,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			race, err := GetRace(db, tt.raceName)
			assert.Equal(t, tt.expectedRaceName, race.Name)
			assert.Equal(t, tt.expectedModStrength, race.ModStrength)
			assert.Equal(t, tt.expectedModDexterity, race.ModDexterity)
			assert.Equal(t, tt.expectedModIntelligence, race.ModIntelligence)
			assert.Equal(t, tt.expectedModWisdom, race.ModWisdom)
			assert.Equal(t, tt.expectError, err)
		})
	}
}
