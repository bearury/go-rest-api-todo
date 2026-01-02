package handler

import (
	"net/http"

	"github.com/bearury/go-rest-api-postgres-todo/entitys"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (handler *Handler) signUp(c *gin.Context) {

	var input entitys.User

	if err := c.BindJSON(&input); err != nil {
		logrus.Errorf("Error binding JSON: %s", err)
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := handler.services.AuthorizationService.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (handler *Handler) signIn(c *gin.Context) {

}
