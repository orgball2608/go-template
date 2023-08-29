package handler

import (
	errorConstants "elearning/error"
	logPkg "elearning/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"

	"elearning/usecase/auth"
)

func Login(c *gin.Context, authUseCase auth.AuthUseCase) {
	var input auth.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		errorCode := errorConstants.CreateErrCodeValidate(c, &input, err)
		c.JSON(errorCode.HTTPCode, errorCode)
		return
	}
	output, errCode := authUseCase.Login(c, &input)
	if errCode != nil {
		c.JSON(errCode.HTTPCode, errCode)
		return
	}
	c.JSON(http.StatusCreated, output)
}

func RefreshToken(c *gin.Context, authUseCase auth.AuthUseCase) {
	var input auth.RefreshTokenInput
	if err := c.ShouldBind(&input); err != nil {
		errorCode := errorConstants.CreateErrCodeValidate(c, &input, err)
		c.JSON(errorCode.HTTPCode, errorCode)
		return
	}
	output, errCode := authUseCase.RefreshToken(c, &input)
	if errCode != nil {
		c.JSON(errCode.HTTPCode, errCode)
		return
	}
	c.JSON(http.StatusOK, output)
}

func ChangePassword(c *gin.Context, authUseCase auth.AuthUseCase) {
	var input auth.ChangePasswordInput
	userIdRaw, _ := c.Get("userId")
	input.UserId = userIdRaw.(int)
	if err := c.ShouldBindJSON(&input); err != nil {
		errorCode := errorConstants.CreateErrCodeValidate(c, &input, err)
		c.JSON(errorCode.HTTPCode, errorCode)
		return
	}
	errCode := authUseCase.ChangePassword(c, &input)
	if errCode != nil {
		c.JSON(errCode.HTTPCode, errCode)
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func RequestResetPassword(c *gin.Context, authUseCase auth.AuthUseCase) {
	var input auth.RequestResetPasswordInput
	if err := c.ShouldBindJSON(&input); err != nil {
		errorCode := errorConstants.CreateErrCodeValidate(c, &input, err)
		c.JSON(errorCode.HTTPCode, errorCode)
		return
	}
	if errCode := authUseCase.RequestResetPassword(c, &input); errCode != nil {
		c.JSON(errCode.HTTPCode, errCode)
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func ResetPassword(c *gin.Context, authUseCase auth.AuthUseCase) {
	var input auth.ResetPasswordInput
	if err := c.ShouldBindJSON(&input); err != nil {
		errorCode := errorConstants.CreateErrCodeValidate(c, &input, err)
		c.JSON(errorCode.HTTPCode, errorCode)
		return
	}
	errCode := authUseCase.ResetPassword(c, &input)
	if errCode != nil {
		c.JSON(errCode.HTTPCode, errCode)
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func CheckTokenResetPassword(c *gin.Context, authUseCase auth.AuthUseCase) {
	var input auth.TokenResetPasswordInput
	token := c.Query("token")
	if token == "" {
		errCode := logPkg.Log(
			c,
			"",
			"",
			&errorConstants.ErrLogicRecordNotFound,
		)
		c.JSON(errCode.HTTPCode, errCode)
		return
	}
	input.Token = token
	errCode := authUseCase.CheckTokenResetPassword(c, &input)
	if errCode != nil {
		c.JSON(errCode.HTTPCode, errCode)
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
