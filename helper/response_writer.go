package helper

import "go-takemikazuchi-api/web"

func WriteSuccess(message string, data interface{}) web.ResponseContract {
	return web.ResponseContract{
		Status:  true,
		Message: message,
		Data:    &data,
	}
}
