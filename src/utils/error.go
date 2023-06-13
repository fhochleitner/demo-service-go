package utils

import (
	"gepaplexx/demo-service/logger"
)

func CheckIfError(err error) {
	if err != nil {
		logger.Fatal(err.Error())
		panic(err)
	}
}
