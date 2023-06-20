package services

import (
	"gepaplexx-demos/demo-service-go/commons"
	"gepaplexx-demos/demo-service-go/logger"
	"gepaplexx-demos/demo-service-go/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JokesHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("Jokes endpoint called")
		c.Writer.WriteHeader(http.StatusOK)
		if c.Request.RequestURI == "/jokes/random" {
			joke := model.RandomJoke()
			_, err := c.Writer.Write([]byte(joke.String()))
			commons.CheckIfError(err)
		}
		// todo handle rate/get/add joke

		c.Next()
	}
}
