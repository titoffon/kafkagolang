package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
	k "github.com/titoffon/kafkagolang/internal/kafka"
)

var address = []string{"localhost:9091", "localhost:9092", "localhost:9093"}
const (
	topic = "my-topic"
)

func main(){
	p, err := k.NewProducer(address)
	if err != nil{
		logrus.Fatal(err)
	}

	for i := 0; i < 100; i++{
		msg := fmt.Sprintf("Kafka message %d", i)
		if err = p.Produce(msg, topic); err != nil {
			logrus.Error(err)
		}
	}
}