package main

import (
	"flag"
	"net/http"
	"time"

	ginzerolog "github.com/dn365/gin-zerolog"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"github.com/d-kv/backend-travel-app/cmd/place-service/config"
	"github.com/d-kv/backend-travel-app/internal/adapter/gateway/oauth_provider/tinkoff"
	ggonicv0 "github.com/d-kv/backend-travel-app/internal/adapter/handler/rest/igin/v0"
	iplace_ctrl_v0 "github.com/d-kv/backend-travel-app/internal/app/controller/v0/place"
	iuser_ctrl_v0 "github.com/d-kv/backend-travel-app/internal/app/controller/v0/user"
	redistoken "github.com/d-kv/backend-travel-app/internal/infra/cache/token/iredis"
	"github.com/d-kv/backend-travel-app/internal/infra/imongo"
	"github.com/d-kv/backend-travel-app/internal/infra/iredis"
	mongoplace "github.com/d-kv/backend-travel-app/internal/infra/repository/place/imongo"
	mongouser "github.com/d-kv/backend-travel-app/internal/infra/repository/user/imongo"
)

const (
	connTimeout          = 5 * time.Second
	oauthProviderTimeout = 3 * time.Second
)

func main() {
	var cfgName string
	flag.StringVar(&cfgName, "config", "docker", "config to run")
	flag.Parse()

	log.Info().
		Str("config", cfgName)

	cfg, err := config.New("cmd/place-service", cfgName)
	if err != nil {
		log.Panic().
			Err(err)
	}

	mongoCl, err := imongo.NewClient(cfg.Storage.Mongo.URI, connTimeout)
	if err != nil {
		log.Panic().
			Err(err)
	}

	redisCl, err := iredis.NewClient(cfg.Cache.Redis.URI, connTimeout)
	if err != nil {
		log.Panic().
			Err(err)
	}

	userRepo := mongouser.NewUserStore(
		mongoCl.
			Database(cfg.Storage.Mongo.DB).
			Collection(cfg.Storage.Mongo.Coll.User),
	)
	placeRepo := mongoplace.NewPlaceStore(
		mongoCl.
			Database(cfg.Storage.Mongo.DB).
			Collection(cfg.Storage.Mongo.Coll.Place),
	)

	tokenCache := redistoken.NewTokenCache(
		redisCl,
	)

	httpCl := &http.Client{
		Timeout: oauthProviderTimeout,
	}

	oauthGateway := tinkoff.New(
		cfg.OAuthProvider.Tinkoff.ID,
		cfg.OAuthProvider.Tinkoff.Secret,
		httpCl,
	)

	placeCtrl := iplace_ctrl_v0.New(
		placeRepo,
	)

	userCtrl := iuser_ctrl_v0.New(
		userRepo,
		tokenCache,
		oauthGateway,
	)

	g := gin.New()
	g.Use(gin.Recovery())
	g.Use(ginzerolog.Logger("gin"))

	restSrv := ggonicv0.New(
		userCtrl,
		placeCtrl,
		g,
	)

	_ = restSrv.Run(cfg.Server.REST.IP, cfg.Server.REST.Port)
}
