package middlewares

import (
	"elearning/config"
	errorConstants "elearning/error"

	logPkg "elearning/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (m *middleware) Authentication(c *gin.Context) {
	headerAuthorization := c.GetHeader("Authorization")
	var token string
	_, err := m.stringService.Sscanf(headerAuthorization, "Bearer %s", &token)
	if err != nil {
		errCode := logPkg.Log(c, "", "", &errorConstants.ErrAuthenticationTokenInvalid)
		c.AbortWithStatusJSON(errCode.HTTPCode, errCode)
		return
	}
	if token == "" {
		errCode := logPkg.Log(c, "", "", &errorConstants.ErrAuthenticationTokenInvalid)
		c.AbortWithStatusJSON(errCode.HTTPCode, errCode)
		return
	}
	verifiedToken, err := m.jwtService.ValidateAccessToken(token)
	if err != nil {
		errCodeRaw := &errorConstants.ErrAuthenticationTokenInvalid
		if err.Error() == errorConstants.ErrTokenExpired.Error() {
			errCodeRaw = &errorConstants.ErrAuthenticationTokenExpired
		}
		errCode := logPkg.Log(c, "", "", errCodeRaw)
		c.AbortWithStatusJSON(errCode.HTTPCode, errCode)
		return
	}

	user, err := m.userRepository.GetById(verifiedToken.UserId)
	if err != nil || user.Status != config.StatusActive {
		errCode := logPkg.Log(c, "", "", &errorConstants.ErrAuthenticationTokenInvalid)
		c.AbortWithStatusJSON(errCode.HTTPCode, errCode)
		return
	}
	role := user.Type

	//TODO implement more role
	var roleCtx string
	switch role {
	case config.USER:
		roleCtx = config.USER
	case config.TEACHER:
		roleCtx = config.TEACHER
	default:
		roleCtx = config.USER
	}
	c.Set("userId", verifiedToken.UserId)
	c.Set("email", verifiedToken.Email)
	c.Set("role", roleCtx)
	loggerRaw, _ := c.Get("logger")
	logger := loggerRaw.(*logrus.Entry)
	logger = logger.WithField("userId", verifiedToken.UserId)
	c.Set("logger", logger)
	errCode := m.Authorization(c)
	if errCode != nil {
		c.AbortWithStatusJSON(errCode.HTTPCode, errCode)
		return
	}
}
