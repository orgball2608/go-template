package auth

import (
	"elearning/config"
	"elearning/domain/entity"
	errorConstants "elearning/error"
	logPkg "elearning/pkg/logger"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RequestResetPasswordInput struct {
	Email string `json:"email" binding:"required,email,customEmail,max=255"`
}

func (au *authUseCase) RequestResetPassword(ctx *gin.Context, input *RequestResetPasswordInput) *errorConstants.ErrorCode {
	user, err := au.userRepository.GetByEmail(input.Email)
	if err != nil {
		logPkg.Log(
			ctx,
			config.LogLevelWarn,
			fmt.Sprintf("Fail to get token for email:%s, err:%s", input.Email, err.Error()),
			&errorConstants.ErrorCode{},
		)
		return nil
	}
	userInfo, err := au.userRepository.GetById(user.Id)
	if err != nil {
		return logPkg.Log(
			ctx,
			config.LogLevelWarn,
			fmt.Sprintf("Fail to get user by email:%s, err:%s", input.Email, err.Error()),
			nil,
		)
	}
	if userInfo.Status == "無効" {
		return logPkg.Log(
			ctx,
			config.LogLevelWarn,
			fmt.Sprintf("Fail to send email is inactivated: %s", input.Email),
			nil,
		)
	}
	userToken, err := au.userTokenRepository.GetByUserId(user.Id)
	if err == nil {
		if userToken.ResendTime.After(time.Now().UTC()) {
			return logPkg.Log(
				ctx,
				config.LogLevelWarn,
				fmt.Sprintf("Fail to reset password again for email %s. ", input.Email),
				&errorConstants.ErrLogicRequestResetPasswordExisted,
			)
		}
		err := au.userTokenRepository.Delete(userToken.Id)
		if err != nil {
			logPkg.Log(
				ctx,
				config.LogLevelWarn,
				fmt.Sprintf("Fail to delete reset password token, err:%s", err.Error()),
				nil,
			)
		}
	}
	token := uuid.New().String()
	location := time.FixedZone("JST", 9*60*60)
	expireAt := time.Now().In(location).Add(time.Hour *
		time.Duration(au.cfg.ResetPassWordDurationHours),
	)
	if _, err = au.userTokenRepository.Create(&entity.UserToken{
		UserId:           user.Id,
		Token:            token,
		ExpireAt:         expireAt.UTC(),
		ResendTime:       time.Now().Add(time.Second * time.Duration(au.cfg.ResendRequestResetPasswordDurationSeconds)).UTC(),
		RecordCreateTime: time.Now(),
	}); err != nil {
		return logPkg.Log(
			ctx,
			config.LogLevelWarn,
			fmt.Sprintf("Fail to create reset password token, err:%s", err.Error()),
			&errorConstants.ErrGeneralSomethingWentWrong,
		)
	}

	// TODO: SEND MAIL

	return nil
}
