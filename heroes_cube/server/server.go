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
	Port string
	DB   *gorm.DB
	App  *fiber.App
}

func (s *Server) setupRoutes() {

	api := s.App.Group("/api/v1")

	api.Get("/", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Heroes Cube API is running!", "status": "ok"})
	})

}

func (s *Server) setupMiddlewares() {
	s.App.Use(recover.New())
	s.App.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
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

	server := &Server{
		Port: config.ServerPort,
		DB:   db,
		App:  fiber.New(),
	}

	return server, nil
}
