package services

import (
	"gepaplexx-demos/demo-service-go/commons"
	"gepaplexx-demos/demo-service-go/logger"
	"gepaplexx-demos/demo-service-go/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AllJokesHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("Jokes endpoint called")

		jokes := model.AllJokes()
		for _, joke := range jokes {
			_, err := c.Writer.Write([]byte(joke.String()))
			commons.CheckIfError(err)
		}
		c.Writer.WriteHeader(http.StatusOK)
		c.Next()

		if c.Request.RequestURI == "/jokes/get/:id" {
			if id, nok := strconv.Atoi(c.Param("id")); nok != nil {
				c.Writer.WriteHeader(http.StatusBadRequest)
				_, err := c.Writer.Write([]byte("Invalid joke id"))
				commons.CheckIfError(err)
				c.Next()
			} else {
				joke := model.GetJoke(id)
				_, err := c.Writer.Write([]byte(joke.String()))
				commons.CheckIfError(err)
				c.Writer.WriteHeader(http.StatusOK)
				c.Next()
			}
		}

		c.Next()
	}
}

func RandomJokeHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		joke := model.RandomJoke()
		_, err := c.Writer.Write([]byte(joke.String()))
		commons.CheckIfError(err)
		c.Writer.WriteHeader(http.StatusOK)
		c.Next()
	}
}

func GetJokeHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if id, nok := strconv.Atoi(c.Param("id")); nok == nil {
			if id > 0 && id <= len(model.AllJokes()) {
				joke := model.GetJoke(id)
				_, err := c.Writer.Write([]byte(joke.String()))
				commons.CheckIfError(err)
				c.Writer.WriteHeader(http.StatusOK)
				c.Next()
			} else {
				handleInvalidJokeId(c)
			}
		} else {
			handleInvalidJokeId(c)
		}
	}
}

func handleInvalidJokeId(c *gin.Context) {
	c.Writer.WriteHeader(http.StatusBadRequest)
	_, err := c.Writer.Write([]byte("Invalid joke id"))
	commons.CheckIfError(err)
	c.Next()
}
