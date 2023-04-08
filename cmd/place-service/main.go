package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"github.com/d-kv/backend-travel-app/cmd/place-service/config"
	resthandlerv0 "github.com/d-kv/backend-travel-app/internal/adapter/handler/rest/v0"
	controllerv0 "github.com/d-kv/backend-travel-app/internal/app/controller/v0"
	"github.com/d-kv/backend-travel-app/internal/infra/repository/mongo"
)

const (
	connTimeout = 3
)

func main() {
	cfg, err := config.New("cmd/place-service", "dev")
	if err != nil {
		log.Panic().Msgf("%s", err)
	}

	mClnt, err := mongo.NewClient(cfg.DB.Mongo.URI, connTimeout*time.Second)
	if err != nil {
		log.Panic().Msgf("%s", err)
	}

	uStore := mongo.NewUserStore(
		mClnt.
			Database(cfg.DB.Mongo.DB).
			Collection(cfg.DB.Mongo.Coll.User),
	)
	pStore := mongo.NewPlaceStore(
		mClnt.
			Database(cfg.DB.Mongo.DB).
			Collection(cfg.DB.Mongo.Coll.Place),
	)

	ctrl := controllerv0.New(pStore, uStore)
	g := gin.Default()
	httpSrv := resthandlerv0.New(ctrl, g)
	_ = httpSrv.Run(cfg.Server.REST.IP, cfg.Server.REST.Port)
}
