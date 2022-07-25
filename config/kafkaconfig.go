package config

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

type KafkaService struct {
	Conn *kafka.Conn
}

const (
	Topic         = "go_kafka_tran"
	BrokerAddress = "localhost:9092"
	Partition     = 0
)

func (k *KafkaService) NewKafkaService() *KafkaService {
	conn, _ := kafka.DialLeader(context.Background(), "tcp", BrokerAddress, Topic, Partition)
	k.Conn = conn
	return k
}

func (k *KafkaService) Produce(message interface{}) {
	uuid, _ := uuid.NewRandom()
	msg, _ := json.Marshal(message)
	k.Conn.WriteMessages(kafka.Message{
		Key:   []byte(uuid[:]),
		Value: msg,
	})
	fmt.Println(message)
	k.Conn.Close()
}
