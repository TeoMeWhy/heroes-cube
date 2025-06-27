package server

import (
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

	api.Get("/", HealthCheckHandler)
	api.Get("/status", HealthCheckHandler)

	api.Get("/races", func(c fiber.Ctx) error {
		races, err := s.Controller.GetRaces()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Falha ao buscar raças"})
		}
		return c.JSON(races)
	})

	api.Get("/classes", func(c fiber.Ctx) error {
		classes, err := s.Controller.GetClasses()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Falha ao buscar classes"})
		}
		return c.JSON(classes)
	})

}

func (s *Server) Start() {
	s.setupMiddlewares()
	s.setupRoutes()
	s.App.Listen(":" + s.Port)
}

func NewServer(config *configs.Config) (*Server, error) {

	db, err := db.GetMySqlClient(*config)
	if err != nil {
		return nil, err
	}

	controller := &Controller{
		Db: db,
	}

	// Initialize the Fiber app
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
