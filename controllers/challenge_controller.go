package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"hello-starwars/services"
	"net/http"
	"strconv"
)

func SayHello(c *gin.Context) {
	c.String(http.StatusOK, "hello world")
}

func HelloStarwars(c *gin.Context) {
	param := c.Query("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.New("param must be integer"))
	}
	person, err := services.ChallengeService.HelloStarwars(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, person)
}
