package handler

import (
	"net/http"

	"github.com/bearury/go-rest-api-postgres-todo/entitys"
	"github.com/gin-gonic/gin"
)

func (handler *Handler) createList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input entitys.Todo
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := handler.services.TodoList.CreateList(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (handler *Handler) getAllList(c *gin.Context) {
	_, ok := c.Get(userCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return
	}

	var input entitys.Todo

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

}

func (handler *Handler) getListById(c *gin.Context) {

}

func (handler *Handler) updateList(c *gin.Context) {

}

func (handler *Handler) deleteList(c *gin.Context) {

}
