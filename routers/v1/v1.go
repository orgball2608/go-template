package v1

import (
	"elearning/middlewares"
	"elearning/usecase/auth"
	"github.com/gin-gonic/gin"
)

func InitV1Router(
	r *gin.RouterGroup,
	middleware middlewares.Middleware,
	authUseCase auth.AuthUseCase,
) {
	r.Use()
	{
		InitAuthRouter(r.Group("/auth"), middleware, authUseCase)
	}
}
