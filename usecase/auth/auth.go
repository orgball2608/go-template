package auth

import (
	"elearning/config"
	"elearning/domain/repository"
	errorConstants "elearning/error"
	stringPkg "elearning/pkg/hasher"
	jwtPkg "elearning/pkg/jwt"

	"github.com/gin-gonic/gin"
)

type AuthUseCase interface {
	Login(ctx *gin.Context, input *LoginInput) (*LoginOutput, *errorConstants.ErrorCode)
	ChangePassword(ctx *gin.Context, input *ChangePasswordInput) *errorConstants.ErrorCode
	CheckTokenResetPassword(ctx *gin.Context, input *TokenResetPasswordInput) *errorConstants.ErrorCode
	RequestResetPassword(ctx *gin.Context, input *RequestResetPasswordInput) *errorConstants.ErrorCode
	ResetPassword(ctx *gin.Context, input *ResetPasswordInput) *errorConstants.ErrorCode
	RefreshToken(ctx *gin.Context, input *RefreshTokenInput) (*RefreshTokenOutput, *errorConstants.ErrorCode)
}

type authUseCase struct {
	cfg                 config.Environment
	jwtService          jwtPkg.JwtService
	stringService       stringPkg.StringService
	userRepository      repository.UserRepository
	cacheRepository     repository.CacheRepository
	userTokenRepository repository.UserTokenRepository
	queueRepository     repository.QueueRepository
}

func NewAuthUseCase(
	cfg config.Environment,
	jwtService jwtPkg.JwtService,
	stringService stringPkg.StringService,
	userRepository repository.UserRepository,
	cacheRepository repository.CacheRepository,
	userTokenRepository repository.UserTokenRepository,
	queueRepository repository.QueueRepository,

) AuthUseCase {
	return &authUseCase{
		cfg:                 cfg,
		userRepository:      userRepository,
		jwtService:          jwtService,
		stringService:       stringService,
		cacheRepository:     cacheRepository,
		userTokenRepository: userTokenRepository,
		queueRepository:     queueRepository,
	}
}
