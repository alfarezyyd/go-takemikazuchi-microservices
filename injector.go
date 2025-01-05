//go:build wireinject
// +build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	universalTranslator "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"go-takemikazuchi-api/category"
	"go-takemikazuchi-api/config"
	"go-takemikazuchi-api/job"
	"go-takemikazuchi-api/routes"
	"go-takemikazuchi-api/user"
	"gorm.io/gorm"
)

var routeSet = wire.NewSet(
	ProvideAuthenticationRoutes,
	ProvideProtectedRoutes,
)

func ProvideAuthenticationRoutes(routerGroup *gin.RouterGroup, userController user.Controller) *routes.AuthenticationRoutes {
	authenticationRoutes := routes.NewAuthenticationRoutes(routerGroup, userController)
	authenticationRoutes.Setup()
	return authenticationRoutes
}

func ProvideProtectedRoutes(routerGroup *gin.RouterGroup, categoryController category.Controller) *routes.ProtectedRoutes {
	protectedRoutes := routes.NewProtectedRoutes(routerGroup, categoryController)
	protectedRoutes.Setup()
	return protectedRoutes
}

var userSet = wire.NewSet(
	user.NewRepository,
	wire.Bind(new(user.Repository), new(*user.RepositoryImpl)),
	user.NewService,
	wire.Bind(new(user.Service), new(*user.ServiceImpl)),
	user.NewHandler,
	wire.Bind(new(user.Controller), new(*user.Handler)),
)

var categorySet = wire.NewSet(
	category.NewRepository,
	wire.Bind(new(category.Repository), new(*category.RepositoryImpl)),
	category.NewService,
	wire.Bind(new(category.Service), new(*category.ServiceImpl)),
	category.NewHandler,
	wire.Bind(new(category.Controller), new(*category.Handler)),
)

var jobSet = wire.NewSet(
	job.NewRepository,
	wire.Bind(new(job.Repository), new(*job.RepositoryImpl)),
	job.NewService,
	wire.Bind(new(job.Service), new(*job.ServiceImpl)),
	job.NewHandler,
	wire.Bind(new(job.Controller), new(*job.Handler)),
)

// wire.go
func InitializeRoutes(
	ginRouterGroup *gin.RouterGroup,
	dbConnection *gorm.DB,
	validatorInstance *validator.Validate,
	engTranslator universalTranslator.Translator,
	viperConfig *viper.Viper,
	mailerService *config.MailerService,
	identityProvider *config.IdentityProvider,
) (*routes.ApplicationRoutes, error) {
	wire.Build(
		wire.Struct(new(routes.ApplicationRoutes), "*"),
		routeSet,
		userSet,
		categorySet,
	)
	return nil, nil
}
