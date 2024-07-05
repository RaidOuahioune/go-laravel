package controllers

import (
	"context"
	"encoding/json"
	"log"

	"demo.com/hello/core/http/utlis"
	"demo.com/hello/db"
	"demo.com/hello/kafka"
	"demo.com/hello/kafka/producers"
	"demo.com/hello/models"
	"demo.com/hello/websockets"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TodoController struct{}

func (t *TodoController) Produce(ctx *gin.Context) {

	var db *gorm.DB = (&db.Database{}).GetInstance()
	var todo models.NewTodo

	// Bind the request body to the user model
	if !utlis.ValidateAndBind(ctx, &todo) {
		return
	}

	var dbTodo = models.Todo{
		Text:   todo.Text,
		Done:   false,
		UserID: 26,
	}

	db.Model(&models.Todo{}).Create(&dbTodo)

	var todoProducer = &producers.TodoProducer{}

	todoProducer.Produce(&dbTodo, ctx)
	ctx.JSON(200, gin.H{
		"message": "Todo has been produced",
	})
}

func (t *TodoController) Index(ctx *gin.Context) {
	conn, _ := (&websockets.WebSocketServer{}).Upgrade(ctx.Writer, ctx.Request)
	reader := (&kafka.KafkaCore{}).NewReader("todos")

	defer conn.Close()
	for {
		message, err := reader.ReadMessage(context.Background())
		var todo models.Todo
		if err != nil {

			log.Fatal("failed to fetch message", err)

		} else {
			json.Unmarshal(message.Value, &todo)
			conn.WriteJSON(todo.Text)

		}
	}

}
