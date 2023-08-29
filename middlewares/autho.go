package middlewares

import (
	"elearning/config"
	errorConstants "elearning/error"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	logPkg "elearning/pkg/logger"

	"github.com/gin-gonic/gin"
)

func (m *middleware) Authorization(c *gin.Context) *errorConstants.ErrorCode {
	jsonPath := "json/master_access_api.json"
	jsonFile, err := os.Open(jsonPath)
	if err != nil {
		errCode := logPkg.Log(
			c,
			config.LogLevelError,
			fmt.Sprintf("Fail to read file with path '%s', err:%s", jsonPath, err.Error()),
			&errorConstants.ErrGeneralSomethingWentWrong,
		)
		return errCode
	}
	defer jsonFile.Close()
	jsonByte, _ := io.ReadAll(jsonFile)
	masterAccessAPI := map[string][]string{}
	_ = json.Unmarshal(jsonByte, &masterAccessAPI)
	method := c.Request.Method
	path := c.FullPath()
	path = strings.Replace(path, ":id", "{id}", -1)
	rolesAccess := masterAccessAPI[fmt.Sprintf("%s %s", method, path)]
	if len(rolesAccess) == 0 {
		return nil
	}
	var isAccess bool
	userRole := c.Value("role")
	for _, role := range rolesAccess {
		if role == "*" {
			isAccess = true
			break
		} else {
			if role == userRole {
				isAccess = true
				break
			}
		}
	}
	if !isAccess {
		errCode := logPkg.Log(c, "", "", &errorConstants.ErrGeneralForbidden)
		return errCode
	}
	return nil
}
