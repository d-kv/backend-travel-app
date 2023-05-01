package config

import (
	"github.com/rs/zerolog/log"

	"github.com/d-kv/backend-travel-app/pkg/app/config"
)

type (
	Storage struct {
		Mongo struct {
			URI  string
			DB   string
			Coll struct {
				Place string
				User  string
			}
		}
	}

	Cache struct {
		Redis struct {
			URI string
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
	Storage       Storage
	Cache         Cache
	OAuthProvider OAuthProvider
	Server        Server
}

func New(path, name string) (*Config, error) {
	var cfg Config
	log.Info().
		Str("config path", path).
		Str("config name", name).
		Msg("trying to load config")

	err := config.Load(&cfg, path, name)
	if err != nil {
		log.Error().
			Err(err)
		return nil, err
	}

	return &cfg, nil
}
