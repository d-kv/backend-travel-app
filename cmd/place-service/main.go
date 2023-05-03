package main

import (
	"flag"
	"net/http"
	"time"

	"github.com/d-kv/backend-travel-app/cmd/place-service/config"
	"github.com/d-kv/backend-travel-app/internal/pkg/infra/imongo"
	"github.com/d-kv/backend-travel-app/internal/pkg/infra/iredis"
	iginplacev0 "github.com/d-kv/backend-travel-app/internal/place_service/adapter/handler/rest/igin/v0"
	ictrlplacev0 "github.com/d-kv/backend-travel-app/internal/place_service/app/controller/v0"
	imongoplace "github.com/d-kv/backend-travel-app/internal/place_service/infra/repository/imongo"
	"github.com/d-kv/backend-travel-app/internal/user_service/adapter/gateway/oauth_provider/tinkoff"
	iginuserv0 "github.com/d-kv/backend-travel-app/internal/user_service/adapter/handler/rest/igin/v0"
	ictrluserv0 "github.com/d-kv/backend-travel-app/internal/user_service/app/controller/v0"
	iredistoken "github.com/d-kv/backend-travel-app/internal/user_service/infra/cache/iredis"
	imongouser "github.com/d-kv/backend-travel-app/internal/user_service/infra/repository/imongo"
	ginzerolog "github.com/dn365/gin-zerolog"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
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

	userRepo := imongouser.New(
		mongoCl.
			Database(cfg.Storage.Mongo.DB).
			Collection(cfg.Storage.Mongo.Coll.User),
	)
	placeRepo := imongoplace.New(
		mongoCl.
			Database(cfg.Storage.Mongo.DB).
			Collection(cfg.Storage.Mongo.Coll.Place),
	)

	tokenCache := iredistoken.New(
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

	placeCtrl := ictrlplacev0.New(
		placeRepo,
	)

	userCtrl := ictrluserv0.New(
		userRepo,
		tokenCache,
		oauthGateway,
	)

	g := gin.New()
	g.Use(gin.Recovery())
	g.Use(ginzerolog.Logger("gin"))

	_ = iginuserv0.New(
		userCtrl,
		g,
	)

	_ = iginplacev0.New(
		placeCtrl,
		g,
	)

	g.Run(cfg.Server.REST.IP + ":" + cfg.Server.REST.Port)
}
