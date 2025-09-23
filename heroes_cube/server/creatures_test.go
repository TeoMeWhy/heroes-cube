package server

import (
	"bytes"
	"encoding/json"
	"heroes_cube/models"
	"net/http"
	"testing"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
)

func TestGetCreature(t *testing.T) {
	db, err := mockDB("TestGetCreature")
	if err != nil {
		t.Fatalf("Failed to mock database: %v", err)
	}

	server, err := mockServer(db)
	if err != nil {
		t.Fatalf("Failed to mock server: %v", err)
	}

	go server.Start()
	time.Sleep(100 * time.Millisecond)

	client := &http.Client{}

	// Sucesso
	resp, err := client.Get("http://localhost:8080/api/v1/creatures/1")
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	var creature models.Creature
	if err := json.NewDecoder(resp.Body).Decode(&creature); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "1", creature.Id)
	assert.Equal(t, "C01", creature.Name)
	assert.Equal(t, "Humano", creature.RaceName)
	assert.Equal(t, "Mago", creature.ClassName)
	assert.Len(t, creature.Inventory.Items, 3)

	// Falha - criatura não encontrada
	resp, err = client.Get("http://localhost:8080/api/v1/creatures/0")
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)

}

func TestPutCreatureSuccess(t *testing.T) {
	db, err := mockDB("TestPutCreatureSuccess")
	if err != nil {
		t.Fatalf("Failed to mock database: %v", err)
	}

	server, err := mockServer(db)
	if err != nil {
		t.Fatalf("Failed to mock server: %v", err)
	}

	go server.Start()
	time.Sleep(100 * time.Millisecond)

	pairTests := []struct {
		name           string
		payload        PayloadPostCreature
		expectedStatus int
		expectedResult models.Creature
		expectError    error
	}{
		{
			name: "Create Creature Success - 02",
			payload: PayloadPostCreature{
				ID:    "2",
				Name:  "C02",
				Race:  "Humano",
				Class: "Mago",
			},
			expectedStatus: http.StatusCreated,
			expectedResult: models.Creature{
				Id:        "2",
				Name:      "C02",
				RaceName:  "Humano",
				ClassName: "Mago",
			},
			expectError: nil,
		},
		{
			name: "Create Creature Success - 03",
			payload: PayloadPostCreature{
				ID:    "3",
				Name:  "C03",
				Race:  "Elfo",
				Class: "Guerreiro",
			},
			expectedStatus: http.StatusCreated,
			expectedResult: models.Creature{
				Id:        "3",
				Name:      "C03",
				RaceName:  "Elfo",
				ClassName: "Guerreiro",
			},
			expectError: nil,
		},
	}

	uri := "http://localhost:8080/api/v1/creatures"

	for _, tt := range pairTests {
		t.Run(tt.name, func(t *testing.T) {
			client := &http.Client{}

			payload, err := json.Marshal(tt.payload)
			if err != nil {
				t.Fatalf("Failed to marshal payload: %v", err)
			}

			resp, err := client.Post(uri, "application/json", bytes.NewBuffer(payload))
			if err != nil {
				t.Fatalf("Failed to make request: %v", err)
			}
			defer resp.Body.Close()

			creature := models.Creature{}
			if err := json.NewDecoder(resp.Body).Decode(&creature); err != nil {
				t.Fatalf("Failed to decode response: %v", err)
			}

			assert.Equal(t, tt.expectedStatus, resp.StatusCode)
			assert.Equal(t, tt.expectedResult.Id, creature.Id)
			assert.Equal(t, tt.expectedResult.Name, creature.Name)
			assert.Equal(t, tt.expectedResult.RaceName, creature.RaceName)
			assert.Equal(t, tt.expectedResult.ClassName, creature.ClassName)
		})
	}

}

func TestPutCreatureFail(t *testing.T) {
	db, err := mockDB("TestPutCreatureFail")
	if err != nil {
		t.Fatalf("Failed to mock database: %v", err)
	}

	server, err := mockServer(db)
	if err != nil {
		t.Fatalf("Failed to mock server: %v", err)
	}

	go server.Start()
	time.Sleep(300 * time.Millisecond)

	pairTests := []struct {
		name           string
		payload        PayloadPostCreature
		expectedStatus int
		expectedResult fiber.Map
		expectError    error
	}{
		{
			name: "Create Creature Already Exists",
			payload: PayloadPostCreature{
				ID:    "1",
				Name:  "C01",
				Race:  "Elfo",
				Class: "Guerreiro",
			},
			expectedStatus: http.StatusBadRequest,
			expectedResult: fiber.Map{"error": "Criatura já existe"},
			expectError:    nil,
		},
		{
			name: "Create Creature with bad race",
			payload: PayloadPostCreature{
				ID:    "4",
				Name:  "C01",
				Race:  "aaaaa",
				Class: "Guerreiro",
			},
			expectedStatus: http.StatusBadRequest,
			expectedResult: fiber.Map{"error": "Raça não encontrada"},
			expectError:    nil,
		},
		{
			name: "Create Creature with bad class",
			payload: PayloadPostCreature{
				ID:    "4",
				Name:  "C01",
				Race:  "Humano",
				Class: "aaaaa",
			},
			expectedStatus: http.StatusBadRequest,
			expectedResult: fiber.Map{"error": "Classe não encontrada"},
			expectError:    nil,
		},
		{
			name: "Create Creature with no ID",
			payload: PayloadPostCreature{
				Name:  "C01",
				Race:  "Humano",
				Class: "Mago",
			},
			expectedStatus: http.StatusBadRequest,
			expectedResult: fiber.Map{
				"error":   "Dados inválidos",
				"details": "Key: 'PayloadPostCreature.ID' Error:Field validation for 'ID' failed on the 'required' tag",
			},
			expectError: nil,
		},
		{
			name: "Create Creature with no name",
			payload: PayloadPostCreature{
				ID:    "4",
				Race:  "Humano",
				Class: "Mago",
			},
			expectedStatus: http.StatusBadRequest,
			expectedResult: fiber.Map{"error": "Dados inválidos", "details": "Key: 'PayloadPostCreature.Name' Error:Field validation for 'Name' failed on the 'required' tag"},
			expectError:    nil,
		},
		{
			name: "Create Creature with no race",
			payload: PayloadPostCreature{
				ID:    "4",
				Name:  "C02",
				Class: "Mago",
			},
			expectedStatus: http.StatusBadRequest,
			expectedResult: fiber.Map{"error": "Dados inválidos", "details": "Key: 'PayloadPostCreature.Race' Error:Field validation for 'Race' failed on the 'required' tag"},
			expectError:    nil,
		},
	}

	uri := "http://localhost:8080/api/v1/creatures"

	for _, tt := range pairTests {
		t.Run(tt.name, func(t *testing.T) {
			client := &http.Client{}

			payload, err := json.Marshal(tt.payload)
			if err != nil {
				t.Fatalf("Failed to marshal payload: %v", err)
			}

			resp, err := client.Post(uri, "application/json", bytes.NewBuffer(payload))
			if err != nil {
				t.Fatalf("Failed to make request: %v", err)
			}
			defer resp.Body.Close()

			respBody := fiber.Map{}
			if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
				t.Fatalf("Failed to decode response: %v", err)
			}

			assert.Equal(t, tt.expectedStatus, resp.StatusCode)
			assert.Equal(t, tt.expectedResult, respBody)

		})
	}

}
