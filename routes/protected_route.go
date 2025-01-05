package routes

import (
	"github.com/gin-gonic/gin"
	"go-takemikazuchi-api/category"
)

type ProtectedRoutes struct {
	routerGroup        *gin.RouterGroup
	categoryController category.Controller
}

func NewProtectedRoutes(routerGroup *gin.RouterGroup, categoryController category.Controller) *ProtectedRoutes {
	return &ProtectedRoutes{
		routerGroup:        routerGroup,
		categoryController: categoryController,
	}
}

func (routerGroup *ProtectedRoutes) Setup() {
}
