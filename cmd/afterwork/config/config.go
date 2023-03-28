package config

import (
	"github.com/d-kv/backend-travel-app/pkg/app/config"
)

type (
	DB struct {
		Mongo struct {
			URI            string
			DBName         string
			CollectionName struct {
				Place string
			}
		}
	}

	Server struct {
		RPC struct {
			Port int
		}
	}
)

type Config struct {
	DB     DB
	Server Server
}

func New(path, name string) (*Config, error) {
	var cfg Config

	err := config.Load(&cfg, path, name)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
