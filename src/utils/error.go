package utils

import (
	"gepaplexx/demo-service/logger"
)

func CheckIfError(err error) {
	if err != nil {
		logger.Error(err.Error())
	}
}
