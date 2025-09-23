package models

import (
	"heroes_cube/clients/db"
	"heroes_cube/configs"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImportClasses(t *testing.T) {

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
		ClassName            string
		expectedClassName    string
		expectedStrength     int
		expectedDexterity    int
		expectedIntelligence int
		expectedWisdom       int
		expectError          error
	}{
		{
			name:                 "Importar raça Guerreiro",
			ClassName:            "Guerreiro",
			expectedClassName:    "Guerreiro",
			expectedStrength:     5,
			expectedDexterity:    3,
			expectedIntelligence: 2,
			expectedWisdom:       2,
			expectError:          nil,
		},
		{
			name:                 "Importar raça Ladinho",
			ClassName:            "Ladino",
			expectedClassName:    "Ladino",
			expectedStrength:     2,
			expectedDexterity:    6,
			expectedIntelligence: 2,
			expectedWisdom:       2,
			expectError:          nil,
		},
		{
			name:                 "Importar raça Mago",
			ClassName:            "Mago",
			expectedClassName:    "Mago",
			expectedStrength:     1,
			expectedDexterity:    3,
			expectedIntelligence: 6,
			expectedWisdom:       2,
			expectError:          nil,
		},
		{
			name:                 "Importar raça Clerigo",
			ClassName:            "Clérigo",
			expectedClassName:    "Clérigo",
			expectedStrength:     1,
			expectedDexterity:    2,
			expectedIntelligence: 3,
			expectedWisdom:       6,
			expectError:          nil,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			Class, err := GetClass(db, tt.ClassName)
			assert.Equal(t, tt.expectedClassName, Class.Name)
			assert.Equal(t, tt.expectedStrength, Class.InitialStrength)
			assert.Equal(t, tt.expectedDexterity, Class.InitialDexterity)
			assert.Equal(t, tt.expectedIntelligence, Class.InitialIntelligence)
			assert.Equal(t, tt.expectedWisdom, Class.InitialWisdom)
			assert.Equal(t, tt.expectError, err)
		})
	}
}
