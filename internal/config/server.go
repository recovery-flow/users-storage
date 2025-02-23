package config

import (
	"github.com/recovery-flow/cifra-rabbit"
	"github.com/recovery-flow/tokens"
	"github.com/recovery-flow/users-storage/internal/service/infra/data/cloud"
	"github.com/recovery-flow/users-storage/internal/service/infra/data/nosql"
	"github.com/sirupsen/logrus"
)

type Service struct {
	Config       *Config
	MongoDB      *nosql.Repo
	Cloud        *cloud.Repo
	Logger       *logrus.Logger
	TokenManager *tokens.TokenManager
	Broker       *cifra_rabbit.Broker
}

func NewService(cfg *Config) (*Service, error) {
	logger := SetupLogger(cfg.Logging.Level, cfg.Logging.Format)
	MongoDb, err := nosql.NewRepositoryNoSql(cfg.Mongo.URI, cfg.Mongo.DbName)
	if err != nil {
		return nil, err
	}
	Storage, err := cloud.NewRepositoryCloud(cfg.Cloud.CloudName, cfg.Cloud.APIKey, cfg.Cloud.APISecret)
	if err != nil {
		return nil, err
	}

	broker, err := cifra_rabbit.NewBroker(cfg.Rabbit.URL, cfg.Rabbit.Exchange)
	if err != nil {
		return nil, err
	}

	tm := tokens.NewTokenManager(cfg.Database.Redis.Addr, cfg.Database.Redis.Password, cfg.Database.Redis.DB, logger, cfg.JWT.AccessToken.TokenLifetime)

	return &Service{
		Config:       cfg,
		MongoDB:      MongoDb,
		Cloud:        Storage,
		Logger:       logger,
		Broker:       broker,
		TokenManager: &tm,
	}, nil
}
