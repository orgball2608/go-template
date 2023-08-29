package middlewares

import (
	"elearning/config"
	"elearning/domain/repository"
	stringPkg "elearning/pkg/hasher"
	jwtPkg "elearning/pkg/jwt"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Middleware interface {
	RestLogger(context *gin.Context)
	Authentication(context *gin.Context)
}

type middleware struct {
	jwtService     jwtPkg.JwtService
	stringService  stringPkg.StringService
	logger         *logrus.Entry
	userRepository repository.UserRepository
	cfg            config.Environment
}

func NewMiddleware(
	jwtService jwtPkg.JwtService,
	stringService stringPkg.StringService,
	logger *logrus.Entry,
	userRepository repository.UserRepository,
	cfg config.Environment,
) Middleware {
	return &middleware{
		jwtService:     jwtService,
		stringService:  stringService,
		logger:         logger,
		userRepository: userRepository,
		cfg:            cfg,
	}
}
