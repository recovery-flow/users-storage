package config

import (
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/recovery-flow/cifra-rabbit"
	"github.com/recovery-flow/tokens"
	"github.com/recovery-flow/users-storage/internal/data/nosql"
	"github.com/sirupsen/logrus"
)

const (
	SERVER = "server"
)

type Service struct {
	Config       *Config
	MongoDB      *nosql.Repo
	Logger       *logrus.Logger
	TokenManager *tokens.TokenManager
	Storage      *cloudinary.Cloudinary
	Broker       *cifra_rabbit.Broker
}

func NewServer(cfg *Config) (*Service, error) {
	logger := SetupLogger(cfg.Logging.Level, cfg.Logging.Format)
	logger.Infof("Mongo %s || %s", cfg.Mongo.URI, cfg.Mongo.Database)
	MongoDb, err := nosql.NewRepositoryNoSql(cfg.Mongo.URI, cfg.Mongo.Database)
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
		MongoDB:      MongoDb,
		Logger:       logger,
		TokenManager: &TokenManager,
		Storage:      Storage,
		Broker:       broker,
	}, nil
}
