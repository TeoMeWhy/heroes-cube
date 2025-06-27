package server

import (
	"fmt"
	"heroes_cube/clients/db"
	"heroes_cube/configs"

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

	server := &Server{
		Port:       config.ServerPort,
		DB:         db,
		App:        fiberApp,
		Controller: controller,
	}

	return server, nil
}
