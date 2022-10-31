package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jahngeor/avito-tech/internal/service"
)

type Controller struct {
	services *service.Service
}

func NewController(services *service.Service) *Controller {
	return &Controller{services: services}
}

func (ctrl *Controller) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		api.POST("")
	}

	return router
}
