package handler

import (
	"fmt"
	"net/http"

	"github.com/bearury/go-rest-api-postgres-todo/entitys"
	"github.com/gin-gonic/gin"
)

func (handler *Handler) createItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	listId := c.Param("id")

	var input entitys.TodoItem
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := handler.services.TodoItem.CreateItem(userId, listId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (handler *Handler) getAllItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	listId := c.Param("id")

	lists, err := handler.services.TodoItem.GetAllItems(userId, listId)
	if err != nil {
		if fmt.Sprint(err) == "item already exists" {
			newErrorResponse(c, http.StatusBadRequest, "Список с таким ID не существует")
			return
		} else {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"items": lists})
}

func (handler *Handler) getItemById(c *gin.Context) {

}

func (handler *Handler) updateItem(c *gin.Context) {

}

func (handler *Handler) deleteItem(c *gin.Context) {

}
