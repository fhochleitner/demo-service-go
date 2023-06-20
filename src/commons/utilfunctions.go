package commons

import "gepaplexx-demos/demo-service-go/logger"

func CheckIfError(err error) {
	if err != nil {
		logger.Error(err.Error())
	}
}
