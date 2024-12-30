package config

import (
	"github.com/cifra-city/tokens"
	"github.com/cifra-city/users-storage/internal/data/db"
	"github.com/sirupsen/logrus"
)

const (
	SERVICE = "service"
)

type Service struct {
	Config       *Config
	Databaser    *db.Databaser
	Logger       *logrus.Logger
	TokenManager *tokens.TokenManager
}

func NewServer(cfg *Config) (*Service, error) {
	logger := SetupLogger(cfg.Logging.Level, cfg.Logging.Format)
	queries, err := db.NewDatabaser(cfg.Database.URL)
	TokenManager := tokens.NewTokenManager(cfg.Redis.Addr, cfg.Redis.Password, cfg.Redis.DB, logger, cfg.JWT.AccessToken.TokenLifetime)

	if err != nil {
		return nil, err
	}

	return &Service{
		Config:       cfg,
		Databaser:    queries,
		Logger:       logger,
		TokenManager: TokenManager,
	}, nil
}
