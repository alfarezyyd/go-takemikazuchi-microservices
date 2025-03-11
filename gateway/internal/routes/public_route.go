package routes

import (
	"github.com/alfarezyyd/go-takemikazuchi-microservices/gateway/internal/handler"
	"github.com/gin-gonic/gin"
)

type PublicRoutes struct {
	routerGroup           *gin.RouterGroup
	transactionController handler.TransactionController
}

func NewPublicRoutes(routerGroup *gin.RouterGroup, transactionController handler.TransactionController) *PublicRoutes {
	return &PublicRoutes{routerGroup: routerGroup.Group("public"), transactionController: transactionController}
}

func (publicRoutes *PublicRoutes) Setup() {
	publicRoutes.routerGroup.POST("/transactions/notifications", publicRoutes.transactionController.Notification)
}
