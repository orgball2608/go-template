package v1

import (
	"elearning/handler"
	"elearning/middlewares"
	"elearning/usecase/auth"

	"github.com/gin-gonic/gin"
)

func InitAuthRouter(
	r gin.IRouter,
	middleware middlewares.Middleware,
	authUseCase auth.AuthUseCase,
) {
	r.POST("/login", func(context *gin.Context) {
		handler.Login(context, authUseCase)
	})
	r.GET("/refresh-token", func(context *gin.Context) {
		handler.RefreshToken(context, authUseCase)
	})
	r.PUT("/change-password", middleware.Authentication, func(context *gin.Context) {
		handler.ChangePassword(context, authUseCase)
	})
	r.PUT("/reset-password", func(context *gin.Context) {
		handler.ResetPassword(context, authUseCase)
	})
	r.GET("/reset-password", func(context *gin.Context) {
		handler.CheckTokenResetPassword(context, authUseCase)
	})
	r.POST("/reset-password", func(context *gin.Context) {
		handler.RequestResetPassword(context, authUseCase)
	})
}
