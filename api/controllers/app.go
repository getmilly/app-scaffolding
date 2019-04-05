package controllers

import (
	"net/http"
	"time"

	"github.com/getmilly/grok/api"
	"github.com/getmilly/grok/nats"
	"github.com/gin-gonic/gin"
)

//AppController ...
type AppController struct {
	producer *nats.Producer
}

//NewAppController ...
func NewAppController(producer *nats.Producer) *AppController {
	return &AppController{producer}
}

//RegisterRoutes ...
func (controller *AppController) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/app", func(c *gin.Context) {
		if err := controller.producer.Publish("app-subject", nats.NewMessage(gin.H{
			"now": time.Now().String(),
		})); err != nil {
			api.ResolveError(c, err)
			return
		}

		c.Status(http.StatusOK)
	})
}
