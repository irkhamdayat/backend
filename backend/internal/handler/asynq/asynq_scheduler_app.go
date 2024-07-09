package asynq

import (
	"github.com/hibiken/asynq"
)

func RegisterSchedulerTask(scheduler *asynq.Scheduler) {
	//defaultAsynqOpts := []asynq.Option{
	//	asynq.Retention(config.Env.Asynq.Retention),
	//	asynq.MaxRetry(config.Env.Asynq.MaxRetry),
	//}

	//_, err := scheduler.Register(cronSpecEvery12Am, asynq.NewTask(task.AsynqCronBoilerplateTask, nil, defaultAsynqOpts...))
	//util.ContinueOrFatal(err)
}
