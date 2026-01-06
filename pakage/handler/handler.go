package handler

import (
	"github.com/bearury/go-rest-api-postgres-todo/pakage/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (handler *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", handler.signUp)
		auth.POST("/sign-in", handler.signIn)
	}

	api := router.Group("/api", handler.userIdentity)
	{
		lists := api.Group("/lists")
		{
			lists.GET("", handler.getAllList)
			lists.POST("", handler.createList)
			lists.GET("/:id", handler.getListById)
			lists.PUT("/:id", handler.updateList)
			lists.DELETE("/:id", handler.deleteList)

			items := lists.Group("/:id/items")
			{
				items.GET("", handler.getAllItem)
				items.POST("", handler.createItem)
				items.GET("/:item_id", handler.getItemById)
				items.PUT("/:item_id", handler.updateItem)
				items.DELETE("/:item_id", handler.deleteItem)
			}
		}

		item := api.Group("/item")
		{
			item.GET("/:item_id", handler.getItemById)
			item.PUT("/:item_id", handler.updateItem)
			item.DELETE("/:item_id", handler.deleteItem)
		}
	}
	return router
}
