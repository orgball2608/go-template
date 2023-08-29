package middlewares

import (
	"bytes"
	"elearning/config"
	"fmt"
	"math"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

type DataWriteLog struct {
	Context *gin.Context
	fields  log.Fields
}

var ExcludingApis = map[string]bool{
	"/api/v1/auth/login":           true,
	"/api/v1/auth/reset-password":  true,
	"/api/v1/auth/refresh-token":   true,
	"/api/v1/auth/change-password": true,
}

var (
	processId = uuid.New().String()
)

func (m *middleware) RestLogger(context *gin.Context) {
	logger := m.logger
	context.Set("logger", logger)
	context.Set("processId", processId)
	start := time.Now()
	blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: context.Writer}
	context.Writer = blw
	context.Next()
	// Log Response
	stop := time.Since(start)
	WriteLog(logger, DataWriteLog{
		Context: context,
		fields: log.Fields{
			"logType":      config.LogTypeResponse,
			"latency":      fmt.Sprintf("%dms", int(math.Ceil(float64(stop.Nanoseconds())/1000000.0))), // time to process
			"responseBody": strings.Replace(blw.body.String(), `"`, `'`, -1),
			"statusCode":   context.Writer.Status(),
		},
	})
}

func WriteLog(logger *log.Entry, writeLog DataWriteLog) {
	var fields = writeLog.fields
	fields["processId"] = processId
	fields["clientIp"] = writeLog.Context.ClientIP()
	fields["method"] = writeLog.Context.Request.Method
	fields["path"] = writeLog.Context.Request.URL.Path
	if query := writeLog.Context.Request.URL.RawQuery; query != "" {
		fields["query"] = query
	}
	var response = fields["responseBody"]
	if ExcludingApis[writeLog.Context.Request.URL.Path] {
		delete(fields, "requestBody")
		delete(fields, "responseBody")
	}
	logger = logger.WithFields(fields)
	statusCode := writeLog.Context.Writer.Status()
	if len(writeLog.Context.Errors) > 0 {
		logger.Error(writeLog.Context.Errors.ByType(gin.ErrorTypePrivate).String())
	} else {
		msg := ""
		if statusCode >= http.StatusBadRequest {
			if response != nil {
				logger = logger.WithFields(log.Fields{
					"responseError": response.(string),
				})
			}
			delete(fields, "responseBody")
			logger.Error(msg)
		} else {
			logger.Info(msg)
		}
	}
}
