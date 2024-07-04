package kafka

// to produce messages

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/segmentio/kafka-go"
)

type KafkaCore struct{}

func (m *KafkaCore) Connect(topic string, partition int) *kafka.Conn {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	kafkaTimeoutStr := os.Getenv("KAFKA_TIMEOUT")
	kafkaTimeout, err := strconv.Atoi(kafkaTimeoutStr)
	if err != nil {
		kafkaTimeout = 10
	}

	conn, err := kafka.DialLeader(context.Background(), "tcp", os.Getenv("KAFKA_ADDR"), topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}
	conn.SetWriteDeadline(time.Now().Add(time.Duration(kafkaTimeout) * time.Second))
	return conn

}
