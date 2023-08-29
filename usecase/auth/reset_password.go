package auth

import (
	"elearning/config"
	errorConstants "elearning/error"
	logPkg "elearning/pkg/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"

	"gorm.io/gorm"
)

type ResetPasswordInput struct {
	Password string `json:"password" binding:"required,customPassword"`
	Token    string `json:"token" binding:"required"`
}

func (au *authUseCase) ResetPassword(ctx *gin.Context, input *ResetPasswordInput) *errorConstants.ErrorCode {
	token, err := au.userTokenRepository.GetByToken(input.Token)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return logPkg.Log(
				ctx,
				config.LogLevelError,
				fmt.Sprintf("Fail to get reset password token, err:%s", err.Error()),
				&errorConstants.ErrGeneralSomethingWentWrong,
			)
		}
		return logPkg.Log(
			ctx,
			config.LogLevelError,
			fmt.Sprintf("Token Invalid:%s, err:%s", input.Token, err.Error()),
			&errorConstants.ErrLogicTokenInvalid,
		)
	}
	if time.Now().After(token.ExpireAt) {
		return logPkg.Log(
			ctx,
			config.LogLevelError,
			fmt.Sprintf("Token Expired:%s, err:%s", input.Token, err.Error()),
			&errorConstants.ErrLogicTokenExpired,
		)
	}
	newPassword, err := au.stringService.HashPassword(input.Password)
	if err != nil {
		return logPkg.Log(
			ctx,
			config.LogLevelError,
			fmt.Sprintf("Hash password fail:%s, err:%s", input.Token, err.Error()),
			&errorConstants.ErrGeneralSomethingWentWrong,
		)
	}
	if err := au.userRepository.ChangePassword(token.UserId, newPassword); err != nil {
		return logPkg.Log(
			ctx,
			config.LogLevelError,
			fmt.Sprintf("Change password fail:%s, err:%s", input.Token, err.Error()),
			&errorConstants.ErrGeneralSomethingWentWrong,
		)
	}
	if err := au.userTokenRepository.Delete(token.Id); err != nil {
		return logPkg.Log(
			ctx,
			config.LogLevelError,
			fmt.Sprintf("Delete token fail:%s, err:%s", input.Token, err.Error()),
			&errorConstants.ErrGeneralSomethingWentWrong,
		)

	}
	user, err := au.userRepository.GetById(token.UserId)
	loginFailKey := fmt.Sprintf("loginFail-%s", user.Email)
	loginFailedCountString, err := au.cacheRepository.Get(ctx, loginFailKey)
	loginFailedCount := 0
	if err == nil {
		loginFailedCount, err = au.stringService.ConvertStringToInt(loginFailedCountString)
	}
	if loginFailedCount > 0 {
		err := au.cacheRepository.Del(ctx, loginFailKey)
		if err != nil {
			logPkg.Log(
				ctx,
				config.LogLevelError,
				fmt.Sprintf("Fail to del key from db, err:%s", err.Error()),
				&errorConstants.ErrGeneralSomethingWentWrong,
			)
		}
	}
	return nil
}
