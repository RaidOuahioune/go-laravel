package consumers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"demo.com/hello/kafka"
	"demo.com/hello/models"
	"github.com/gin-gonic/gin"
)

type TodoConsumer struct{}

func (m *TodoConsumer) Consume(webSockerHandler func(ctx *gin.Context)) {

	reader := (&kafka.KafkaCore{}).NewReader("todos")
	for {
		message, err := reader.FetchMessage(context.TODO())
		if err != nil {
			log.Fatal("failed to fetch message", err)
		}
		var t models.Todo
		json.Unmarshal(message.Value, &t)

		fmt.Println("message received: ", t.Text)
		reader.CommitMessages(context.TODO(), message)
	}

}
