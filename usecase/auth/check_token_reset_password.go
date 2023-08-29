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

type TokenResetPasswordInput struct {
	Token string
}

func (au *authUseCase) CheckTokenResetPassword(ctx *gin.Context, input *TokenResetPasswordInput) *errorConstants.ErrorCode {
	token, err := au.userTokenRepository.GetByToken(input.Token)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return logPkg.Log(
				ctx,
				config.LogLevelError,
				fmt.Sprintf("Fail to get user token:%s, err:%s", input.Token, err.Error()),
				&errorConstants.ErrGeneralSomethingWentWrong,
			)
		}
		return logPkg.Log(
			ctx,
			config.LogLevelError,
			fmt.Sprintf("Fail to get user token:%s, err:%s", input.Token, err.Error()),
			&errorConstants.ErrLogicTokenInvalid,
		)
	}
	if time.Now().UTC().After(token.ExpireAt.UTC()) {
		return logPkg.Log(
			ctx,
			config.LogLevelError,
			fmt.Sprintf("Token reset password expired:%s", input.Token),
			&errorConstants.ErrLogicTokenExpired,
		)
	}
	return nil
}
