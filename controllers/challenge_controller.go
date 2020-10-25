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

func HelloStarWars(c *gin.Context) {
	param := c.Query("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.New("param must be integer"))
		return
	}

	person, err2 := services.ChallengeService.SayHelloFromSWAPI(id)
	if err2 != nil {
		c.JSON(err2.CodeStatus(), err2)
		return
	}
	c.JSON(http.StatusOK, person)
}
