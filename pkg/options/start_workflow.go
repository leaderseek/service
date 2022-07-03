package options

import (
	"time"

	"github.com/leaderseek/service/pkg/task_queue"
	"go.temporal.io/api/enums/v1"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/temporal"
)

func StartWorkflowOptionsDefault() client.StartWorkflowOptions {
	return client.StartWorkflowOptions{
		TaskQueue:                                task_queue.Leaderseek,
		WorkflowExecutionTimeout:                 15 * time.Second,
		WorkflowIDReusePolicy:                    enums.WORKFLOW_ID_REUSE_POLICY_REJECT_DUPLICATE,
		WorkflowExecutionErrorWhenAlreadyStarted: true,
		RetryPolicy: &temporal.RetryPolicy{
			MaximumAttempts: 1,
		},
	}
}
