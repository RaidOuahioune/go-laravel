package job

import (
	"log"
	"time"

	"demo.com/hello/core/job/tasks"

	"github.com/hibiken/asynq"
)

func Client() {
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: GetRedisAddress()})
	defer client.Close()
	// an example to regiser a task for a later execution.
	// The task will be processed in 24 hours.
	// The task will be retried at most 10 times
	// the task will be timed out after 3 minutes.
	task, err := tasks.NewWelcomeEmailTask(42) // send email to user with id 42
	if err != nil {
		log.Fatalf("could not create task: %v", err)
	}
	info, err := client.Enqueue(task, asynq.ProcessIn(24*time.Hour), asynq.MaxRetry(10), asynq.Timeout(3*time.Minute))
	if err != nil {
		log.Fatalf("could not schedule task: %v", err)
	}
	log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)

}
