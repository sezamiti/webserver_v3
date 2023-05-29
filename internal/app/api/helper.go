package api

import (
	"github.com/sezamiti/go2/StandardWebServer/storage"
	"github.com/sirupsen/logrus"
)

var (
	prefix string = "/api/v1"
)

// Пытаемся откунфигурировать наш Api инстанс (а конкретнее - поле logger)
func (a *Api) configreLoggerField() error {
	log_level, err := logrus.ParseLevel(a.config.LoggerLevel)
	if err != nil {
		return err
	}
	a.logger.SetLevel(log_level)
	return nil
}

// Пытаемся отконфигурировать маршрутизатор (а конкретнее поле router Api)
func (a *Api) configreRouterField() {
	a.router.HandleFunc(prefix+"/articles", a.GetAllArticles).Methods("GET")
	a.router.HandleFunc(prefix+"/articles/{id}", a.GetArticleById).Methods("GET")
	a.router.HandleFunc(prefix+"/articles/{id}", a.DeleteArticleById).Methods("DELETE")
	a.router.HandleFunc(prefix+"/articles", a.PostArticle).Methods("POST")
	a.router.HandleFunc(prefix+"/user/register", a.PostUserRegister).Methods("POST")

}

// Пытаемся отконфигурировать наше хранилище (storage Api)
func (a *Api) configreStorageField() error {
	storage := storage.New(a.config.Storage)
	//Пытаемся установить соединениение, если невозможно - возвращаем ошибку!
	if err := storage.Open(); err != nil {
		return err
	}
	a.storage = storage
	return nil
}
