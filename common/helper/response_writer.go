package helper

import "github.com/alfarezyyd/go-takemikazuchi-microservices-common/web"

func WriteSuccess(message string, data interface{}) web.ResponseContract {
	return web.ResponseContract{
		Status:  true,
		Message: message,
		Data:    &data,
	}
}
