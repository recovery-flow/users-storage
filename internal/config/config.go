package config

import (
	"os"
	"time"

	_ "github.com/lib/pq" // postgres driver don`t delete
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type ServerConfig struct {
	Name     string `mapstructure:"name"`
	Port     string `mapstructure:"port"`
	BasePath string `mapstructure:"base_path"`
	TestMode bool   `mapstructure:"test_mode"`
	Log      struct {
		Level  string `mapstructure:"level"`
		Format string `mapstructure:"format"`
	} `mapstructure:"logging"`
}

type DatabaseConfig struct {
	SQL struct {
		URL string `mapstructure:"url"`
	} `mapstructure:"sql"`

	Redis struct {
		Addr     string `mapstructure:"addr"`
		Password string `mapstructure:"password"`
		DB       int    `mapstructure:"db"`
		Lifetime int    `mapstructure:"lifetime"`
	} `mapstructure:"redis"`

	Mongo struct {
		URI  string `mapstructure:"uri"`
		Name string `mapstructure:"db_name"`
	}
}

type OAuthConfig struct {
	Google struct {
		ClientID     string `mapstructure:"client_id"`
		ClientSecret string `mapstructure:"client_secret"`
		RedirectURL  string `mapstructure:"redirect_url"`
	}
}

type JWTConfig struct {
	AccessToken struct {
		SecretKey     string        `mapstructure:"secret_key"`
		TokenLifetime time.Duration `mapstructure:"token_lifetime"`
	} `mapstructure:"access_token"`
	RefreshToken struct {
		SecretKey     string        `mapstructure:"secret_key"`
		EncryptionKey string        `mapstructure:"encryption_key"`
		TokenLifetime time.Duration `mapstructure:"token_lifetime"`
	} `mapstructure:"refresh_token"`
	Bin struct {
		Addr     string `mapstructure:"addr"`
		Password string `mapstructure:"password"`
		DB       int    `mapstructure:"db"`
		Lifetime int    `mapstructure:"lifetime"`
	} `mapstructure:"bin"`
}

type RabbitConfig struct {
	URL      string `mapstructure:"url"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

type KafkaConfig struct {
	Brokers      []string      `mapstructure:"brokers"`
	Topic        string        `mapstructure:"topic"`
	GroupID      string        `mapstructure:"group_id"`
	DialTimeout  time.Duration `mapstructure:"dial_timeout"`
	ReadTimeout  time.Duration `mapstructure:"read_timeout"`
	WriteTimeout time.Duration `mapstructure:"write_timeout"`
	RequiredAcks string        `mapstructure:"required_acks"`
}

type SwaggerConfig struct {
	Enabled bool   `mapstructure:"enabled"`
	URL     string `mapstructure:"url"`
	Port    string `mapstructure:"port"`
}

type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	JWT      JWTConfig      `mapstructure:"jwt"`
	OAuth    OAuthConfig    `mapstructure:"oauth"`
	Kafka    KafkaConfig    `mapstructure:"kafka"`
	Rabbit   RabbitConfig   `mapstructure:"rabbit"`
	Database DatabaseConfig `mapstructure:"database"`
	Swagger  SwaggerConfig  `mapstructure:"swagger"`
}

func LoadConfig() (*Config, error) {
	configPath := os.Getenv("KV_VIPER_FILE")
	if configPath == "" {
		configPath = "./../config_local.yaml"
		//return nil, errors.New("KV_VIPER_FILE env var is not set")
	}

	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Errorf("error reading config file: %s", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, errors.Errorf("error unmarshalling config: %s", err)
	}

	return &config, nil
}
