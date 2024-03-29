package main

import (
	"flag"
	"net/http"
	"time"

	ginzerolog "github.com/dn365/gin-zerolog"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"github.com/d-kv/backend-travel-app/cmd/afterwork-backend/config"
	"github.com/d-kv/backend-travel-app/internal/pkg/imongo"
	"github.com/d-kv/backend-travel-app/internal/pkg/iredis"
	ictrlplacev0 "github.com/d-kv/backend-travel-app/internal/place-service/controller/v0"
	iginplacev0 "github.com/d-kv/backend-travel-app/internal/place-service/handler/rest/igin/v0"
	imongoplace "github.com/d-kv/backend-travel-app/internal/place-service/repository/imongo"
	iredistoken "github.com/d-kv/backend-travel-app/internal/user-service/cache/iredis"
	ictrluserv0 "github.com/d-kv/backend-travel-app/internal/user-service/controller/v0"
	"github.com/d-kv/backend-travel-app/internal/user-service/gateway/oauth_provider/tinkoff"
	iginuserv0 "github.com/d-kv/backend-travel-app/internal/user-service/handler/rest/igin/v0"
	imongouser "github.com/d-kv/backend-travel-app/internal/user-service/repository/imongo"
)

const (
	connTimeout          = 5 * time.Second
	oauthProviderTimeout = 3 * time.Second
)

func main() {
	var cfgName string
	flag.StringVar(&cfgName, "config", "docker", "config to run")
	flag.Parse()

	cfg, err := config.New("cmd/afterwork-backend", cfgName)
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

	err = g.Run(cfg.Server.REST.IP + ":" + cfg.Server.REST.Port)
	if err != nil {
		panic(err)
	}
}
