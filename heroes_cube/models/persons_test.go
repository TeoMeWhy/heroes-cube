package models

import (
	"fmt"
	"heroes_cube/clients/db"
	"heroes_cube/configs"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImportPerson(t *testing.T) {

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
		name                 string
		RaceName             string
		ClassName            string
		expectedRaceName     string
		expectedClassName    string
		expectedStrength     int
		expectedDexterity    int
		expectedIntelligence int
		expectedWisdom       int
		expectError          error
	}{
		{
			name:                 "Cria Humano Guerreiro",
			RaceName:             "Humano",
			ClassName:            "Guerreiro",
			expectedRaceName:     "Humano",
			expectedClassName:    "Guerreiro",
			expectedStrength:     5,
			expectedDexterity:    3,
			expectedIntelligence: 2,
			expectedWisdom:       2,
			expectError:          nil,
		},
		{
			name:                 "Importar Humano Ladinho",
			RaceName:             "Humano",
			ClassName:            "Ladino",
			expectedRaceName:     "Humano",
			expectedClassName:    "Ladino",
			expectedStrength:     2,
			expectedDexterity:    6,
			expectedIntelligence: 2,
			expectedWisdom:       2,
			expectError:          nil,
		},
		{
			name:                 "Importar Humano Mago",
			RaceName:             "Humano",
			ClassName:            "Mago",
			expectedRaceName:     "Humano",
			expectedClassName:    "Mago",
			expectedStrength:     1,
			expectedDexterity:    3,
			expectedIntelligence: 6,
			expectedWisdom:       2,
			expectError:          nil,
		},
		{
			name:                 "Importar Humano Clerigo",
			RaceName:             "Humano",
			ClassName:            "Clérigo",
			expectedRaceName:     "Humano",
			expectedClassName:    "Clérigo",
			expectedStrength:     1,
			expectedDexterity:    2,
			expectedIntelligence: 3,
			expectedWisdom:       6,
			expectError:          nil,
		},
	}

	for i, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			person, err := NewPerson(fmt.Sprintf("%d", i), tt.RaceName, tt.ClassName, db)
			assert.Equal(t, tt.expectedClassName, person.ClassName)
			assert.Equal(t, tt.expectedStrength, person.PtsStrength)
			assert.Equal(t, tt.expectedDexterity, person.PtsDexterity)
			assert.Equal(t, tt.expectedIntelligence, person.PtsIntelligence)
			assert.Equal(t, tt.expectedWisdom, person.PtsWisdom)
			assert.Equal(t, tt.expectError, err)
		})
	}
}
