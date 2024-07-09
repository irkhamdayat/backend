package util

import (
	"github.com/goccy/go-json"
	"github.com/hibiken/asynq"
	"github.com/sirupsen/logrus"

	"github.com/Halalins/backend/internal/common/constant"
)

func BindingAsynqPayload(t *asynq.Task, out any) error {
	err := json.Unmarshal(t.Payload(), &out)
	if err != nil {
		return err
	}
	return nil
}

func processPayloadToAsynqTask(task string, body any) (*asynq.Task, error) {
	byteData, err := json.Marshal(&body)
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(task, byteData), nil
}

func enqueueTask(asynqClient *asynq.Client, task *asynq.Task, opts ...asynq.Option) error {
	result, err := asynqClient.Enqueue(task, opts...)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"task":    task.Type(),
			"payload": string(task.Payload()),
		}).Error(err)
		return err
	}

	logrus.WithFields(logrus.Fields{
		"taskID":  result.ID,
		"task:":   task.Type(),
		"payload": string(task.Payload()),
	}).Info("Enqueue Task Success ")

	return nil
}

func ProcessPayloadAndEnqueueTask(asynqClient *asynq.Client, task string, body any, opts ...asynq.Option) error {
	asynqTask, err := processPayloadToAsynqTask(task, body)
	if err != nil {
		return err
	}

	if len(opts) <= 0 {
		opts = []asynq.Option{
			asynq.Retention(constant.AsynqDefaultRetention),
			asynq.MaxRetry(constant.AsynqDefaultMaxRetry),
		}
	}

	return enqueueTask(asynqClient, asynqTask, opts...)
}
