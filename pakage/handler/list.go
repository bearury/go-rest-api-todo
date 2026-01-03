package handler

import (
	"go/types"
	"net/http"

	"github.com/bearury/go-rest-api-postgres-todo/entitys"
	"github.com/gin-gonic/gin"
)

func (handler *Handler) createList(c *gin.Context) {

}

func (handler *Handler) getAllList(c *gin.Context) {
	id, ok := c.Get(userCtx)
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
