package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/d-kv/backend-travel-app/cmd/place-service/config"
	resthandlerv0 "github.com/d-kv/backend-travel-app/internal/adapter/handler/rest/v0"
	controllerv0 "github.com/d-kv/backend-travel-app/internal/app/controller/v0"
	"github.com/d-kv/backend-travel-app/internal/infra/logger/deflog"
	"github.com/d-kv/backend-travel-app/internal/infra/repository/mongo"
)

const (
	connTimeout = 10
)

func main() {
	cfg, err := config.New("cmd/place-service", "dev")
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}

	l := deflog.New(log.Default())
	mClnt, _ := mongo.NewClient(l, cfg.DB.Mongo.URI, connTimeout*time.Second)

	uStore := mongo.NewUserStore(
		l,
		mClnt.
			Database(cfg.DB.Mongo.DB).
			Collection(cfg.DB.Mongo.Coll.User),
	)
	pStore := mongo.NewPlaceStore(
		l,
		mClnt.
			Database(cfg.DB.Mongo.DB).
			Collection(cfg.DB.Mongo.Coll.Place),
	)

	ctrl := controllerv0.New(l, pStore, uStore)
	g := gin.Default()
	httpSrv := resthandlerv0.New(ctrl, g)
	httpSrv.Run(cfg.Server.REST.IP, cfg.Server.REST.Port)
}
