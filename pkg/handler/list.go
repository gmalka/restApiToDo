package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	todo "github.com/gmalka/rest_api"
)

func (h *Handler) createList(c *gin.Context) {
	id, ok := c.Get(userCtx)
	if ok == false {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return
	}

	var input todo.TodoList
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
}

func (h *Handler) getAllLists(c *gin.Context) {

}

func (h *Handler) getListById(c *gin.Context) {

}

func (h *Handler) updateList(c *gin.Context) {

}

func (h *Handler) deleteList(c *gin.Context) {

}