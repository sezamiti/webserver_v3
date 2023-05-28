package api

import (
	"github.com/sezamiti/go2/StandardWebServer/storage"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (a Api) configureLoggerField() error {
	log_level, err := logrus.ParseLevel(a.config.LoggerLevel)
	if err != nil {
		return err
	}
	a.logger.SetLevel(log_level)
	return nil
}

func (a *Api) configureRouterField() {
	a.router.HandleFunc("/", func(w http.ResponseWriter, request *http.Request) {
		w.Write([]byte("hello! This is rest API"))
	})

}

func (a *Api) configreStorageField() error {
	storage := storage.New(a.config.Storage)
	//Пытаемся установить соединениение, если невозможно - возвращаем ошибку!
	if err := storage.Open(); err != nil {
		return err
	}
	a.storage = storage
	return nil
}
