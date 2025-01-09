package config

import (
	"github.com/cifra-city/cifra-rabbit"
	"github.com/cifra-city/tokens"
	"github.com/cifra-city/users-storage/internal/data/sql"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/sirupsen/logrus"
)

const (
	SERVER = "server"
)

type Service struct {
	Config       *Config
	Databaser    *sql.Repo
	Logger       *logrus.Logger
	TokenManager *tokens.TokenManager
	Storage      *cloudinary.Cloudinary
	Broker       *cifra_rabbit.Broker
}

func NewServer(cfg *Config) (*Service, error) {
	logger := SetupLogger(cfg.Logging.Level, cfg.Logging.Format)
	queries, err := sql.NewRepoSQL(cfg.Database.URL)
	if err != nil {
		return nil, err
	}
	TokenManager := tokens.NewTokenManager(cfg.Redis.Addr, cfg.Redis.Password, cfg.Redis.DB, logger, cfg.JWT.AccessToken.TokenLifetime)
	Storage, err := InitCloudinaryClient(*cfg)
	if err != nil {
		return nil, err
	}

	broker, err := cifra_rabbit.NewBroker(cfg.Rabbit.URL, cfg.Rabbit.Exchange)
	if err != nil {
		return nil, err
	}

	return &Service{
		Config:       cfg,
		Databaser:    queries,
		Logger:       logger,
		TokenManager: TokenManager,
		Storage:      Storage,
		Broker:       broker,
	}, nil
}
