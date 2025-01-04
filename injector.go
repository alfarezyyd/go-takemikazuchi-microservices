//go:build wireinject
// +build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"go-takemikazuchi-api/config"
	"go-takemikazuchi-api/routes"
	"go-takemikazuchi-api/user"
	"gorm.io/gorm"
)

var routeSet = wire.NewSet(
	ProvideAuthenticationRoutes,
)

func ProvideAuthenticationRoutes(routerGroup *gin.RouterGroup, userController user.Controller) *routes.AuthenticationRoutes {
	authenticationRoutes := routes.NewAuthenticationRoutes(routerGroup, userController)
	authenticationRoutes.Setup()
	return authenticationRoutes
}

var userSet = wire.NewSet(
	user.NewRepository,
	wire.Bind(new(user.Repository), new(*user.RepositoryImpl)),
	user.NewService,
	wire.Bind(new(user.Service), new(*user.ServiceImpl)),
	user.NewHandler,
	wire.Bind(new(user.Controller), new(*user.Handler)),
)

// wire.go
func InitializeRoutes(
	ginRouterGroup *gin.RouterGroup,
	dbConnection *gorm.DB,
	validatorInstance *validator.Validate,
	engTranslator ut.Translator,
	viperConfig *viper.Viper,
	mailerService *config.MailerService,
	identityProvider *config.IdentityProvider,
) (*routes.ApplicationRoutes, error) {
	wire.Build(
		wire.Struct(new(routes.ApplicationRoutes), "*"),
		routeSet,
		userSet,
	)
	return nil, nil
}
