package auth

import (
	"elearning/config"
	errorConstants "elearning/error"
	jwtPkg "elearning/pkg/jwt"
	"errors"
	"fmt"
	"time"

	logPkg "elearning/pkg/logger"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoginInput struct {
	Email    string `json:"email" binding:"required,email,customEmail,max=255"`
	Password string `json:"password" binding:"required"`
}

type LoginUser struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type LoginOutput struct {
	User         LoginUser `json:"user"`
	AccessToken  string    `json:"accessToken"`
	RefreshToken string    `json:"refreshToken"`
}

func (au *authUseCase) Login(ctx *gin.Context, input *LoginInput) (*LoginOutput, *errorConstants.ErrorCode) {
	loginFailKey := fmt.Sprintf("loginFail-%s", input.Email)
	loginFailedCountString, err := au.cacheRepository.Get(ctx, loginFailKey)
	if err != nil && !errors.Is(err, errorConstants.ErrKeyDoesNotExist) {
		errCode := logPkg.Log(
			ctx,
			config.LogLevelError,
			fmt.Sprintf("Fail to get key '%s' in kvDatabase, err:%s", loginFailKey, err.Error()),
			&errorConstants.ErrGeneralSomethingWentWrong,
		)
		return nil, errCode
	}
	loginFailedCount := 0
	if err == nil {
		loginFailedCount, err = au.stringService.ConvertStringToInt(loginFailedCountString)
		if err != nil {
			errCode := logPkg.Log(
				ctx,
				config.LogLevelError,
				fmt.Sprintf("Fail to convert string '%v' to number, err:%s", loginFailedCountString, err.Error()),
				&errorConstants.ErrGeneralSomethingWentWrong,
			)
			return nil, errCode
		}
		if loginFailedCount >= au.cfg.LoginFailLimit {
			errCode := logPkg.Log(ctx, "", "", &errorConstants.ErrLogicLoginFailedManyTimes)
			return nil, errCode
		}
	}
	user, err := au.userRepository.GetByEmail(input.Email)
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			errCode := logPkg.Log(ctx, "", "", &errorConstants.ErrLogicLoginFailed)
			return nil, errCode
		}
		errCode := logPkg.Log(
			ctx,
			config.LogLevelError,
			fmt.Sprintf("Fail to get user info with email '%s', err:%s", input.Email, err.Error()),
			&errorConstants.ErrGeneralSomethingWentWrong,
		)
		return nil, errCode
	}
	err = au.stringService.CheckHashPassword(user.Password, input.Password)
	if err != nil {
		switch true {
		case loginFailedCount+1 == au.cfg.LoginFailLimit:
			subErr := au.cacheRepository.Set(
				ctx,
				loginFailKey,
				loginFailedCount+1,
				0,
			)
			if subErr != nil {
				logPkg.Log(
					ctx,
					config.LogLevelError,
					fmt.Sprintf("Fail to set key to db, err:%s", err.Error()),
					&errorConstants.ErrGeneralSomethingWentWrong,
				)
			}
		case loginFailedCount == 0:
			subErr := au.cacheRepository.Set(
				ctx,
				loginFailKey,
				loginFailedCount+1,
				time.Minute*time.Duration(au.cfg.LoginFailDurationMinutes),
			)
			if subErr != nil {
				logPkg.Log(
					ctx,
					config.LogLevelError,
					fmt.Sprintf("Fail to set key to db, err:%s", err.Error()),
					&errorConstants.ErrGeneralSomethingWentWrong,
				)
			}
		default:
			subErr := au.cacheRepository.Incr(
				ctx,
				loginFailKey,
			)
			if subErr != nil {
				logPkg.Log(
					ctx,
					config.LogLevelError,
					fmt.Sprintf("Fail to incr key, err:%s", err.Error()),
					&errorConstants.ErrGeneralSomethingWentWrong,
				)
			}
		}
		errCode := logPkg.Log(ctx, "", "", &errorConstants.ErrLogicLoginFailed)
		return nil, errCode
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
	accessToken, err := au.jwtService.GenerateAccessToken(&jwtPkg.GenerateTokenInput{
		UserId: user.Id,
		Email:  user.Email,
	})
	if err != nil {
		errCode := logPkg.Log(
			ctx,
			config.LogLevelError,
			fmt.Sprintf("Fail to generate access token, err:%s", err.Error()),
			&errorConstants.ErrGeneralSomethingWentWrong,
		)
		return nil, errCode
	}
	refreshToken, err := au.jwtService.GenerateRefreshToken(&jwtPkg.GenerateTokenInput{
		UserId: user.Id,
		Email:  user.Email,
	})
	if err != nil {
		errCode := logPkg.Log(
			ctx,
			config.LogLevelError,
			fmt.Sprintf("Fail to generate refresh token, err:%s", err.Error()),
			&errorConstants.ErrGeneralSomethingWentWrong,
		)
		return nil, errCode
	}
	user.LastLoginAt = time.Now().UTC()
	if _, err := au.userRepository.UpdateLastLoginAt(user.Id, user); err != nil {
		return nil, logPkg.Log(
			ctx,
			config.LogLevelError,
			fmt.Sprintf("Fail to update user, err:%s", err.Error()),
			&errorConstants.ErrGeneralSomethingWentWrong,
		)
	}
	return &LoginOutput{
		LoginUser{
			Id:    user.Id,
			Email: user.Email,
			Name:  user.Name,
		},
		accessToken,
		refreshToken,
	}, nil
}
