package tasks

import (
	"encoding/json"

	"github.com/hibiken/asynq"
)

// A list of task types.
const (
	TypeWelcomeEmail = "email:welcome"
)

// Task payload for any email related tasks.
type EmailTaskPayload struct {
	// ID for the email recipient.
	UserID int
}

func NewWelcomeEmailTask(id int) (*asynq.Task, error) {
	payload, err := json.Marshal(EmailTaskPayload{UserID: id})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeWelcomeEmail, payload), nil
}
