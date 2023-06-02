package server

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"

	logger "github.com/hyperversalblocks/trainandflex/pkg/logrus"
	"github.com/hyperversalblocks/trainandflex/pkg/postgres"
)

type server struct {
	config *Config
	logger *logrus.Logger
	db     *pgxpool.Pool
	router *gin.Engine
}

func New() error {
	cfg, err := GetConfig()
	if err != nil {
		return fmt.Errorf("error initiating server: %w", err)
	}

	db, err := postgres.InitPool(postgres.Config{
		Host:     cfg.Database.Host,
		Password: cfg.Database.Password,
		Port:     cfg.Database.Port,
		Name:     cfg.Database.DBName,
		Username: cfg.Database.DBUsername,
	})
	if err != nil {
		return fmt.Errorf("error initiating postgres: %w", err)
	}

	loggerObject, err := logger.Init(logger.Logger{
		LogEnv:   cfg.Logger.LogEnv,
		LogLevel: cfg.Logger.LogLevel,
	})

	serverObject := &server{
		config: cfg,
		logger: loggerObject,
		db:     db,
		router: gin.Default(),
	}

	if err = serverObject.setupRoutes(); err != nil {
		return err
	}

	err = serverObject.router.Run(serverObject.config.Address)
	if err != nil {
		return fmt.Errorf("error running gin server: %w", err)
	}

	return nil
}

func GetConfig() (*Config, error) {
	cfg := Config{}
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, fmt.Errorf("error loading config: %w", err)
	}

	return &cfg, nil
}

func setCORs(router *gin.Engine) {
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))
}

func (s *server) setupRoutes() error {
	// s.router.Use(s.ErrorHandler)

	// publicRoutes := s.router.Group("/")
	// {
	// }
	return nil
}
