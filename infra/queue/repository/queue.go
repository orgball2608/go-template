package repository

import (
	"elearning/config"
	"elearning/domain/repository"
	"elearning/infra/queue"
	logPkg "elearning/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	asynqLib "github.com/hibiken/asynq"
)

type queueRepository struct {
	queueClient *asynq.QueueClient
}

func NewQueueRepository(queueClient *asynq.QueueClient) repository.QueueRepository {
	return &queueRepository{
		queueClient: queueClient,
	}
}

func (qr *queueRepository) Enqueue(queueName string, payload []byte) (string, error) {
	task := asynqLib.NewTask(queueName, payload)
	taskInfo, err := qr.queueClient.Enqueue(task)
	if err != nil {
		return "", err
	}
	return taskInfo.ID, nil
}

func (qr *queueRepository) EnqueueWithLogger(ctx *gin.Context, queueName string, payload []byte) error {
	processId, _ := ctx.Get("processId")
	logger := logPkg.InitLogger(config.LogTypeQueue, processId.(string))
	jobId, err := qr.Enqueue(queueName, payload)
	if err != nil {
		logger.Error(logrus.Fields{"jobId": jobId, "queueName": config.TypeSendEmail, "jobPayload": string(payload), "error": err.Error()}, "")
		return err
	}
	logger.Info(logrus.Fields{"jobId": jobId, "queueName": config.TypeSendEmail, "jobPayload": string(payload)}, "")
	return nil
}

func (qr *queueRepository) EnqueueWithSchedule(queueName string, payload []byte, sendAt time.Time) (string, error) {
	task := asynqLib.NewTask(queueName, payload, asynqLib.ProcessAt(sendAt))
	taskInfo, err := qr.queueClient.Enqueue(task)
	if err != nil {
		return "", err
	}
	return taskInfo.ID, nil
}
