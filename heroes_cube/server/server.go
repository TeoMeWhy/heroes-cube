package server

import (
	"fmt"
	"heroes_cube/clients/db"
	"heroes_cube/configs"
	"heroes_cube/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"

	"gorm.io/gorm"
)

type Server struct {
	Port       string
	DB         *gorm.DB
	App        *fiber.App
	Controller *Controller
	Validate   *validator.Validate
}

func (s *Server) setupMiddlewares() {
	s.App.Use(recover.New())
	s.App.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

}

func (s *Server) setupRoutes() {

	api := s.App.Group("/api/v1")

	// Rota de Health Check
	api.Get("/", HealthCheckHandler)
	api.Get("/status", HealthCheckHandler)

	// Rota para dados de raça
	api.Get("/races", func(c fiber.Ctx) error {
		races, err := s.Controller.GetRaces()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Falha ao buscar raças"})
		}
		return c.JSON(fiber.Map{"races": races, "count": len(races)})

	})

	// Rota para dados de classes
	api.Get("/classes", func(c fiber.Ctx) error {
		classes, err := s.Controller.GetClasses()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Falha ao buscar classes"})
		}
		return c.JSON(fiber.Map{"classes": classes, "count": len(classes)})
	})

	// Rota para dados de itens
	api.Get("/items", func(c fiber.Ctx) error {
		items, err := s.Controller.GetItems()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Falha ao buscar itens"})
		}
		return c.JSON(fiber.Map{"items": items, "count": len(items)})
	})

	// Rota para dados de criaturas
	api.Get("/creatures", func(c fiber.Ctx) error {
		creatures, err := s.Controller.GetCreatures()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Falha ao buscar criaturas"})
		}
		return c.JSON(fiber.Map{"creatures": creatures, "count": len(creatures)})
	})

	// Rota para buscar criatura por ID
	api.Get("/creatures/:id", func(c fiber.Ctx) error {
		id := c.Params("id")
		creature, err := s.Controller.GetCreaturesByID(id)
		if err != nil {

			if err == gorm.ErrRecordNotFound {
				msg := fmt.Sprintf("Criatura com ID %s não encontrada", id)
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": msg})
			}

			msg := fmt.Sprintf("Falha ao buscar criatura com ID %s - %s", id, err.Error())
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": msg})
		}
		return c.JSON(creature)
	})

	// Rota para criar criatura
	api.Post("/creatures", func(c fiber.Ctx) error {

		var payload PayloadPostCreature
		if err := c.Bind().Body(&payload); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Dados inválidos"})
		}

		if err := s.Validate.Struct(&payload); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Dados inválidos", "details": err.Error()})
		}

		creature, err := s.Controller.PostCreature(payload)
		if err != nil {
			if err == models.ErrorCreatureAlreadyExists {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Criatura já existe"})
			} else if err == models.ErrorRaceNotFound {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Raça não encontrada"})
			} else if err == models.ErrorClassNotFound {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Classe não encontrada"})
			}
		}

		return c.Status(fiber.StatusCreated).JSON(creature)

	})

	// Rota para deletar criatura
	api.Delete("/creatures/:id", func(c fiber.Ctx) error {
		id := c.Params("id")
		if err := s.Controller.DeleteCreature(id); err != nil {
			if err == gorm.ErrRecordNotFound {
				msg := fmt.Sprintf("Criatura com ID %s não encontrada", id)
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": msg})
			}
			msg := fmt.Sprintf("Falha ao deletar criatura com ID %s - %s", id, err.Error())
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": msg})
		}
		c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Criatura deletada com sucesso"})
		return nil
	})

	api.Post("/creatures/:id/add_exp", func(c fiber.Ctx) error {

		id := c.Params("id")
		var payload PayloadAddExp
		if err := c.Bind().Body(&payload); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Dados inválidos"})
		}

		if err := s.Validate.Struct(&payload); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Dados inválidos", "details": err.Error()})
		}

		creature, err := s.Controller.AddExpPoints(id, payload.ExpPoins)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				msg := fmt.Sprintf("Criatura com ID %s não encontrada", id)
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": msg})
			}
			msg := fmt.Sprintf("Falha ao adicionar pontos de experiência - %s", err.Error())
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": msg})
		}

		c.Status(fiber.StatusOK).JSON(creature)
		return nil
	})

	api.Post("/creatures/:id/add_skill_points", func(c fiber.Ctx) error {

		id := c.Params("id")
		var payload models.SkillPoints
		if err := c.Bind().Body(&payload); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Dados inválidos"})
		}

		if err := s.Validate.Struct(&payload); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Dados inválidos", "details": err.Error()})
		}

		creature, err := s.Controller.AddSkillPoints(id, payload)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				msg := fmt.Sprintf("Criatura com ID %s não encontrada", id)
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": msg})
			}
			msg := fmt.Sprintf("Falha ao adicionar pontos de habilidade - %s", err.Error())
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": msg})
		}

		c.Status(fiber.StatusOK).JSON(creature)
		return nil
	})

	// Rota criar item no inventário
	api.Post("/inventory/:id_inventory/item", func(c fiber.Ctx) error {

		idInventory := c.Params("id_inventory")

		payload := PayloadPostInventoryItem{}
		if err := c.Bind().Body(&payload); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Dados inválidos"})
		}

		if err := s.Validate.Struct(&payload); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Dados inválidos", "details": err.Error()})
		}

		if err := s.Controller.PostInventoryItem(idInventory, payload.ItemID); err != nil {
			if err == models.ErrorItemIdNotFoundOnInventory {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Item não encontrado no inventário"})
			}
			if err == models.ErrorItemTypeAlreadyExists {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Tipo de item já existe no inventário. Remova o tipo para adicionar novo item"})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Falha ao adicionar item ao inventário"})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Item adicionado ao inventário com sucesso"})
	})

	api.Delete("/inventory/:id_inventory/item/:item_id", func(c fiber.Ctx) error {
		idInventory := c.Params("id_inventory")
		itemID := c.Params("item_id")

		item, err := s.Controller.DeleteInventoryItem(idInventory, itemID)
		if err != nil {
			if err == models.ErrorItemIdNotFoundOnInventory {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Item não encontrado no inventário"})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Falha ao remover item do inventário"})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"item": item})

	})

}

func (s *Server) Start() {
	s.setupMiddlewares()
	s.setupRoutes()
	s.App.Listen(":" + s.Port)
}

func HealthCheckHandler(c fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "ok", "message": "Heroes Cube API is running!"})
}

func NewServer(config *configs.Config) (*Server, error) {

	db, err := db.GetMySqlClient(*config)
	if err != nil {
		return nil, err
	}

	controller := &Controller{
		Db: db,
	}

	fiberApp := fiber.New(fiber.Config{
		AppName: "Heroes Cube API",
	})

	validate := validator.New()

	server := &Server{
		Port:       config.ServerPort,
		DB:         db,
		App:        fiberApp,
		Controller: controller,
		Validate:   validate,
	}

	return server, nil
}
