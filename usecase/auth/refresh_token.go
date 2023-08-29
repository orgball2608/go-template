package auth

import (
	"elearning/config"
	errorConstants "elearning/error"
	jwtPkg "elearning/pkg/jwt"
	"fmt"

	logPkg "elearning/pkg/logger"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RefreshTokenInput struct {
	RefreshToken string `form:"refreshToken" binding:"required"`
}

type RefreshTokenOutput struct {
	AccessToken  string    `json:"accessToken"`
	RefreshToken string    `json:"refreshToken"`
	User         LoginUser `json:"user"`
}

func (au *authUseCase) RefreshToken(ctx *gin.Context, input *RefreshTokenInput) (*RefreshTokenOutput, *errorConstants.ErrorCode) {
	verifiedToken, err := au.jwtService.ValidateRefreshToken(input.RefreshToken)
	if err != nil {
		errCodeRaw := &errorConstants.ErrAuthenticationTokenInvalid
		if err.Error() == errorConstants.ErrTokenExpired.Error() {
			errCodeRaw = &errorConstants.ErrAuthenticationTokenExpired
		}
		return nil, logPkg.Log(
			ctx,
			config.LogLevelWarn,
			fmt.Sprintf("Fail to validate token, err:%s", err.Error()),
			errCodeRaw,
		)
	}
	accessToken, err := au.jwtService.GenerateAccessToken(&jwtPkg.GenerateTokenInput{
		UserId: verifiedToken.UserId,
		Email:  verifiedToken.Email,
	})
	if err != nil {
		return nil, logPkg.Log(
			ctx,
			config.LogLevelError,
			fmt.Sprintf("Fail to generate access token, err:%s", err.Error()),
			&errorConstants.ErrGeneralSomethingWentWrong,
		)
	}
	refreshToken, err := au.jwtService.GenerateRefreshToken(&jwtPkg.GenerateTokenInput{
		UserId: verifiedToken.UserId,
		Email:  verifiedToken.Email,
	})
	if err != nil {
		return nil, logPkg.Log(
			ctx,
			config.LogLevelError,
			fmt.Sprintf("Fail to generate refresh token, err:%s", err.Error()),
			&errorConstants.ErrGeneralSomethingWentWrong,
		)
	}
	if verifiedToken.UserId == -1 {
		return &RefreshTokenOutput{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
			User: LoginUser{
				Id: -1,
			},
		}, nil
	}

	userRaw, err := au.userRepository.GetById(verifiedToken.UserId)
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, logPkg.Log(
				ctx,
				config.LogLevelWarn,
				fmt.Sprintf("User not found with id=%d, err: %s ", verifiedToken.UserId, err.Error()),
				&errorConstants.ErrLogicRecordNotFound,
			)
		}
		return nil, logPkg.Log(
			ctx,
			config.LogLevelError,
			fmt.Sprintf("Fail to get detail user, err:%s", err.Error()),
			&errorConstants.ErrGeneralSomethingWentWrong,
		)
	}
	return &RefreshTokenOutput{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User: LoginUser{
			Id:    userRaw.Id,
			Email: userRaw.Email,
			Name:  userRaw.Name,
		},
	}, nil
}
