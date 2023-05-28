package api

import (
	"github.com/gorilla/mux"
	"github.com/sezamiti/go2/StandardWebServer/storage"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Api struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	//Добавление поля для работы с хранилищем
	storage *storage.Storage
}

func New(config *Config) *Api {
	return &Api{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (api *Api) Start() error {

	if err := api.configureLoggerField(); err != nil {
		return err
	}

	api.logger.Info("starting api server at port", api.config.BindAddr)
	api.configureRouterField()
	return http.ListenAndServe(api.config.BindAddr, api.router)
}
