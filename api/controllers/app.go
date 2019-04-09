package controllers

import (
	"net/http"
	"time"

	"github.com/getmilly/grok/api"
	"github.com/getmilly/grok/mongodb"
	"github.com/getmilly/grok/nats"
	"github.com/gin-gonic/gin"
)

//Applicattion ...
type Applicattion struct {
	Name string `bson:"name"`
}

//AppController ...
type AppController struct {
	producer *nats.Producer
	mongodb  *mongodb.MongoConnection
}

//NewAppController ...
func NewAppController(producer *nats.Producer, mongodb *mongodb.MongoConnection) *AppController {
	return &AppController{
		producer: producer,
		mongodb:  mongodb,
	}
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

	router.POST("/app", func(c *gin.Context) {

		collection := controller.mongodb.Session.DB("app").C("scaffolding")

		err := collection.Insert(Applicattion{
			Name: "scaffolding",
		})

		if err != nil {
			api.ResolveError(c, err)
			return
		}

		c.Status(http.StatusCreated)
	})
}
