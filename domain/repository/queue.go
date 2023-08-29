package repository

import (
	"time"

	"github.com/gin-gonic/gin"
)

type QueueRepository interface {
	Enqueue(queueName string, payload []byte) (string, error)
	EnqueueWithSchedule(queueName string, payload []byte, sendAt time.Time) (string, error)
	EnqueueWithLogger(ctx *gin.Context, queueName string, payload []byte) error
}
