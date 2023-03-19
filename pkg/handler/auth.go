package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	todo "github.com/gmalka/rest_api"
	"github.com/sirupsen/logrus"
)

func (h *Handler) signUp(c *gin.Context) {
	logrus.Println("Starting create user")
	var input todo.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) signIn(c *gin.Context) {

}
