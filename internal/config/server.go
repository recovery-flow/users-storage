package config

import (
	"github.com/cifra-city/tokens"
	"github.com/cifra-city/users-storage/internal/data/db"
	"github.com/cloudinary/cloudinary-go/v2"
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
	Storage      *cloudinary.Cloudinary
}

func NewServer(cfg *Config) (*Service, error) {
	logger := SetupLogger(cfg.Logging.Level, cfg.Logging.Format)
	queries, err := db.NewDatabaser(cfg.Database.URL)
	if err != nil {
		return nil, err
	}
	TokenManager := tokens.NewTokenManager(cfg.Redis.Addr, cfg.Redis.Password, cfg.Redis.DB, logger, cfg.JWT.AccessToken.TokenLifetime)
	Storage, err := InitCloudinaryClient(*cfg)
	if err != nil {
		return nil, err
	}

	return &Service{
		Config:       cfg,
		Databaser:    queries,
		Logger:       logger,
		TokenManager: TokenManager,
		Storage:      Storage,
	}, nil
}
