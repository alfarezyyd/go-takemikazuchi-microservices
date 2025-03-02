package routes

import (
	"github.com/gin-gonic/gin"
	"go-takemikazuchi-microservices/internal/transaction"
)

type PublicRoutes struct {
	routerGroup           *gin.RouterGroup
	transactionController transaction.Controller
}

func NewPublicRoutes(routerGroup *gin.RouterGroup, transactionController transaction.Controller) *PublicRoutes {
	return &PublicRoutes{routerGroup: routerGroup.Group("public"), transactionController: transactionController}
}

func (publicRoutes *PublicRoutes) Setup() {
	publicRoutes.routerGroup.POST("/transactions/notifications", publicRoutes.transactionController.Notification)
}
