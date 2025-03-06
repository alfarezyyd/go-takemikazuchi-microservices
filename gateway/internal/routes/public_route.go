package routes

import (
	"github.com/gin-gonic/gin"
)

type PublicRoutes struct {
}

func NewPublicRoutes(routerGroup *gin.RouterGroup) *PublicRoutes {
	return nil
}

func (publicRoutes *PublicRoutes) Setup() {
}
