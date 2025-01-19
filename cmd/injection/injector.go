//go:build wireinject
// +build wireinject

package injection

import (
	"github.com/gin-gonic/gin"
	universalTranslator "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"go-takemikazuchi-api/configs"
	category2 "go-takemikazuchi-api/internal/category"
	job2 "go-takemikazuchi-api/internal/job"
	job_application2 "go-takemikazuchi-api/internal/job_application"
	routes2 "go-takemikazuchi-api/internal/routes"
	user2 "go-takemikazuchi-api/internal/user"
	"gorm.io/gorm"
)

var routeSet = wire.NewSet(
	ProvideAuthenticationRoutes,
	ProvideProtectedRoutes,
)

func ProvideAuthenticationRoutes(routerGroup *gin.RouterGroup, userController user2.Controller) *routes2.AuthenticationRoutes {
	authenticationRoutes := routes2.NewAuthenticationRoutes(routerGroup, userController)
	authenticationRoutes.Setup()
	return authenticationRoutes
}

func ProvideProtectedRoutes(routerGroup *gin.RouterGroup,
	categoryController category2.Controller,
	jobController job2.Controller,
	viperConfig *viper.Viper) *routes2.ProtectedRoutes {
	protectedRoutes := routes2.NewProtectedRoutes(routerGroup, categoryController, jobController, viperConfig)
	protectedRoutes.Setup()
	return protectedRoutes
}

var userSet = wire.NewSet(
	user2.NewRepository,
	wire.Bind(new(user2.Repository), new(*user2.RepositoryImpl)),
	user2.NewService,
	wire.Bind(new(user2.Service), new(*user2.ServiceImpl)),
	user2.NewHandler,
	wire.Bind(new(user2.Controller), new(*user2.Handler)),
)

var categorySet = wire.NewSet(
	category2.NewRepository,
	wire.Bind(new(category2.Repository), new(*category2.RepositoryImpl)),
	category2.NewService,
	wire.Bind(new(category2.Service), new(*category2.ServiceImpl)),
	category2.NewHandler,
	wire.Bind(new(category2.Controller), new(*category2.Handler)),
)

var jobSet = wire.NewSet(
	job2.NewRepository,
	wire.Bind(new(job2.Repository), new(*job2.RepositoryImpl)),
	job2.NewService,
	wire.Bind(new(job2.Service), new(*job2.ServiceImpl)),
	job2.NewHandler,
	wire.Bind(new(job2.Controller), new(*job2.Handler)),
)

var jobApplicationSet = wire.NewSet(
	job_application2.NewRepository,
	wire.Bind(new(job_application2.Repository), new(*job_application2.RepositoryImpl)),
	job_application2.NewService,
	wire.Bind(new(job_application2.Service), new(*job_application2.ServiceImpl)),
	job_application2.NewHandler,
	wire.Bind(new(job_application2.Controller), new(*job_application2.Handler)),
)

// wire.go
func InitializeRoutes(
	ginRouterGroup *gin.RouterGroup,
	dbConnection *gorm.DB,
	validatorInstance *validator.Validate,
	engTranslator universalTranslator.Translator,
	viperConfig *viper.Viper,
	mailerService *configs.MailerService,
	identityProvider *configs.IdentityProvider,
) (*routes2.ApplicationRoutes, error) {
	wire.Build(
		wire.Struct(new(routes2.ApplicationRoutes), "*"),
		routeSet,
		userSet,
		jobSet,
		categorySet,
	)
	return nil, nil
}
