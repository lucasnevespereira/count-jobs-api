package api

import (
	"count-jobs/api/router"
	"count-jobs/configs"
	"count-jobs/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Api struct {
	router *gin.Engine
	config configs.ApiConfig
}

func New() *Api {
	api := &Api{}
	api.setup()
	return api
}

func (a *Api) setup() {
	a.config = configs.LoadApi()
	a.router = router.Init()
}

func (a *Api) Run() error {
	server := http.Server{
		Addr:    fmt.Sprintf(":%d", a.config.Port),
		Handler: a.router.Handler(),
	}

	err := server.ListenAndServe()
	if err != nil {
		return err
	}

	utils.Logger.Info("Running api")

	return nil
}
