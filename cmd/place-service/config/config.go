package config

import (
	"github.com/d-kv/backend-travel-app/pkg/app/config"
)

type (
	DB struct {
		Mongo struct {
			URI  string
			DB   string
			Coll struct {
				Place string
				User  string
			}
		}
	}

	OAuthProvider struct {
		Tinkoff struct {
			ID     string
			Secret string
		}
	}

	Server struct {
		REST struct {
			IP   string
			Port string
		}
	}
)

type Config struct {
	DB            DB
	OAuthProvider OAuthProvider
	Server        Server
}

func New(path, name string) (*Config, error) {
	var cfg Config

	err := config.Load(&cfg, path, name)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
