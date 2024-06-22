package server

import (
	"context"
	"log"

	"github.com/MohamedDhiaZoghlami/go-commerce/storage"
	"github.com/MohamedDhiaZoghlami/go-commerce/storage/postgres"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// Server server instance
type Server struct {
	Log     *logrus.Logger
	storage storage.Storage
}

func NewServer(logger *logrus.Logger) (*Server, error) {
	if logger == nil {
		logger = logrus.New()
		logger.SetLevel(logrus.InfoLevel)
	}
	db, err := postgres.Open(context.Background(), logger)
	if err != nil {
		log.Fatalf("Error Opening DB connection : %s", err)
	}
	return &Server{
		Log:     logger,
		storage: db,
	}, nil

}

func (s *Server) Run(port string) error {
	r := gin.Default()
	root := r.Group("/")
	{
		root.GET("/ping", func(c *gin.Context) { c.JSON(200, gin.H{"ping": "pong"}) })
	}
	if err := r.Run(port); err != nil {
		return errors.Errorf("serving on %s failed: %v", port, err)
	}
	return nil
}
