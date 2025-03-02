package helper

import "go-takemikazuchi-microservices/web"

func WriteSuccess(message string, data interface{}) web.ResponseContract {
	return web.ResponseContract{
		Status:  true,
		Message: message,
		Data:    &data,
	}
}
