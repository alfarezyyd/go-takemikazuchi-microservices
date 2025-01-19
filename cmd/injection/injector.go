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
	categoryFeature "go-takemikazuchi-api/internal/category"
	jobFeature "go-takemikazuchi-api/internal/job"
	jobApplicationFeature "go-takemikazuchi-api/internal/job_application"
	jobResourceFeature "go-takemikazuchi-api/internal/job_resource"
	"go-takemikazuchi-api/internal/routes"
	"go-takemikazuchi-api/internal/storage"
	userFeature "go-takemikazuchi-api/internal/user"
	"gorm.io/gorm"
)

var routeSet = wire.NewSet(
	ProvideAuthenticationRoutes,
	ProvideProtectedRoutes,
)

func ProvideAuthenticationRoutes(routerGroup *gin.RouterGroup, userController userFeature.Controller) *routes.AuthenticationRoutes {
	authenticationRoutes := routes.NewAuthenticationRoutes(routerGroup, userController)
	authenticationRoutes.Setup()
	return authenticationRoutes
}

func ProvideProtectedRoutes(routerGroup *gin.RouterGroup,
	categoryController categoryFeature.Controller,
	jobController jobFeature.Controller,
	viperConfig *viper.Viper) *routes.ProtectedRoutes {
	protectedRoutes := routes.NewProtectedRoutes(routerGroup, categoryController, jobController, viperConfig)
	protectedRoutes.Setup()
	return protectedRoutes
}

var userSet = wire.NewSet(
	userFeature.NewRepository,
	wire.Bind(new(userFeature.Repository), new(*userFeature.RepositoryImpl)),
	userFeature.NewService,
	wire.Bind(new(userFeature.Service), new(*userFeature.ServiceImpl)),
	userFeature.NewHandler,
	wire.Bind(new(userFeature.Controller), new(*userFeature.Handler)),
)

var categorySet = wire.NewSet(
	categoryFeature.NewRepository,
	wire.Bind(new(categoryFeature.Repository), new(*categoryFeature.RepositoryImpl)),
	categoryFeature.NewService,
	wire.Bind(new(categoryFeature.Service), new(*categoryFeature.ServiceImpl)),
	categoryFeature.NewHandler,
	wire.Bind(new(categoryFeature.Controller), new(*categoryFeature.Handler)),
)

var jobSet = wire.NewSet(
	jobFeature.NewRepository,
	wire.Bind(new(jobFeature.Repository), new(*jobFeature.RepositoryImpl)),
	jobFeature.NewService,
	wire.Bind(new(jobFeature.Service), new(*jobFeature.ServiceImpl)),
	jobFeature.NewHandler,
	wire.Bind(new(jobFeature.Controller), new(*jobFeature.Handler)),
)

var jobApplicationSet = wire.NewSet(
	jobApplicationFeature.NewRepository,
	wire.Bind(new(jobApplicationFeature.Repository), new(*jobApplicationFeature.RepositoryImpl)),
	jobApplicationFeature.NewService,
	wire.Bind(new(jobApplicationFeature.Service), new(*jobApplicationFeature.ServiceImpl)),
	jobApplicationFeature.NewHandler,
	wire.Bind(new(jobApplicationFeature.Controller), new(*jobApplicationFeature.Handler)),
)

var jobResourceSet = wire.NewSet(
	jobResourceFeature.NewRepository,
	wire.Bind(new(jobResourceFeature.Repository), new(*jobResourceFeature.RepositoryImpl)),
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
) (*routes.ApplicationRoutes, error) {
	wire.Build(
		wire.Struct(new(routes.ApplicationRoutes), "*"),
		routeSet,
		userSet,
		jobSet,
		categorySet,
		jobResourceSet,
		storage.ProvideFileStorage, // Fungsi untuk memilih implementasi yang sesuai
	)
	return nil, nil
}
