package config

type LogType string

var (
	LogTypeRequestResponse LogType = "request/response"
	LogTypeUploadFile      LogType = "uploadFile"
	LogTypeExportFile      LogType = "exportFile"
	LogTypeDatabaseQuery   LogType = "databaseQuery"
	LogTypeHandler         LogType = "handler"
	LogTypeQueue           LogType = "queue"
	LogTypeRequest         LogType = "request"
	LogTypeResponse        LogType = "response"
)
