package worker

import (
	"github.com/friendsofgo/errors"
	"github.com/leaderseek/definition/activity/db"
	"github.com/leaderseek/definition/workflow"
	"github.com/leaderseek/service/pkg/app"
	"github.com/leaderseek/service/pkg/config"
	"github.com/leaderseek/service/pkg/task_queue"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.uber.org/zap"
)

type Worker struct {
	logger         *zap.Logger
	temporalWorker worker.Worker
}

func NewWorker(cfg *config.WorkerConfig) (*Worker, error) {
	if err := cfg.Validate(); err != nil {
		return nil, errors.Wrap(err, "failed to validate config")
	}

	logger, err := cfg.Logger.ToZap().Build()
	if err != nil {
		return nil, errors.Wrap(err, "failed to start zap logger")
	}

	c, err := client.Dial(client.Options{
		HostPort: cfg.TemporalClient.HostPort,
		Logger:   app.ZapToTemporalLogger(logger),
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to dial temporal client")
	}

	opts := cfg.TemporalWorkerOptions.ToTemporal()

	w := worker.New(c, task_queue.Leaderseek, *opts)

	return &Worker{logger: logger, temporalWorker: w}, nil
}

func (w *Worker) Run() {
	a := new(db.Config)
	w.temporalWorker.RegisterActivity(a)
	w.temporalWorker.RegisterWorkflow(workflow.TeamCreate)
	w.temporalWorker.RegisterWorkflow(workflow.TeamAddMembers)

	err := w.temporalWorker.Run(worker.InterruptCh())

	w.logger.Fatal("temporal worker stopped running", zap.Error(err))
}
