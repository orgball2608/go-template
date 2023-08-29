package auth

import (
	"elearning/config"
	errorConstants "elearning/error"
	logPkg "elearning/pkg/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

type ChangePasswordInput struct {
	UserId          int
	CurrentPassword string `json:"currentPassword" binding:"required"`
	NewPassword     string `json:"newPassword" binding:"required,customPassword"`
}

func (au *authUseCase) ChangePassword(ctx *gin.Context, input *ChangePasswordInput) *errorConstants.ErrorCode {
	user, err := au.userRepository.GetById(input.UserId)
	if err != nil {
		return logPkg.Log(
			ctx,
			config.LogLevelError,
			fmt.Sprintf("Fail to get user info %d, err:%s", input.UserId, err.Error()),
			&errorConstants.ErrGeneralSomethingWentWrong,
		)
	}
	if err = au.stringService.CheckHashPassword(user.Password, input.CurrentPassword); err != nil {
		return logPkg.Log(
			ctx,
			config.LogLevelError,
			fmt.Sprintf("Cannot check hash password, err:%s", err.Error()),
			&errorConstants.ErrLogicPasswordIncorrect,
		)
	}
	if strings.EqualFold(input.NewPassword, input.CurrentPassword) {
		return logPkg.Log(
			ctx,
			config.LogLevelError,
			fmt.Sprintf("current password is incorrect"),
			&errorConstants.ErrLogicUseCurrentPasswordToChangePassword,
		)
	}
	newPassword, err := au.stringService.HashPassword(input.NewPassword)
	if err != nil {
		return logPkg.Log(
			ctx,
			config.LogLevelError,
			fmt.Sprintf("Fail to encrypt password, err:%s", err.Error()),
			&errorConstants.ErrGeneralSomethingWentWrong,
		)
	}
	if err = au.userRepository.ChangePassword(input.UserId, newPassword); err != nil {
		return logPkg.Log(
			ctx,
			config.LogLevelError,
			fmt.Sprintf("Fail to change password: %d, err:%s", user.Id, err.Error()),
			&errorConstants.ErrGeneralSomethingWentWrong,
		)
	}
	return nil
}
