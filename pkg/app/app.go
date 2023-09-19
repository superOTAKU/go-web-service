package app

import (
	"errors"
	"go-web-service/pkg/cmd"
	"go-web-service/pkg/config"
	"go-web-service/pkg/db"
	"go-web-service/pkg/logging"

	"github.com/gin-gonic/gin"
)

type Application struct {
	config *config.ApplicationConfig
	router *gin.Engine
}

func ListenAndServe(config *config.ApplicationConfig) (*Application, error) {
	var app Application
	if err := logging.InitLogger(&config.Logging); err != nil {
		return nil, errors.New("fail to init app logger: " + err.Error())
	}
	router := gin.Default()
	cmd.PublishCommands(router)
	if err := db.InitDb(&config.Db); err != nil {
		return nil, errors.New("init db error: " + err.Error())
	}
	if err := router.Run(config.Server.GetListenAddr()); err != nil {
		return nil, err
	}
	app.router = router
	app.config = config
	return &app, nil
}
