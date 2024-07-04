package producers

import (
	"encoding/json"

	"demo.com/hello/kafka"
	"demo.com/hello/models"
	"github.com/gin-gonic/gin"
	kafkaGo "github.com/segmentio/kafka-go"
)

type TodoProducer struct {
}

func (m *TodoProducer) Produce(todo *models.Todo, ctx *gin.Context) {

	conn := (&kafka.KafkaCore{}).Connect("todos", 0)
	todoMap := map[string]interface{}{}
	todoMap["text"] = todo.Text
	todoMap["done"] = todo.Done
	todoMap["user_id"] = todo.UserID
	todoMap["id"] = todo.ID
	var todoJson, err = json.Marshal(todoMap)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": "Internal server error",
		})
		return
	}

	conn.WriteMessages(
		kafkaGo.Message{Value: todoJson},
	)
}
