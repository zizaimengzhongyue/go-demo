// kafka-go demo
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

var protocol = "tcp"
var addr = "localhost:9092"
var topic = "new-topic-02"
var partition = 0

func write(done chan bool) {
	conn, err := kafka.DialLeader(context.Background(), protocol, addr, topic, partition)
	defer conn.Close()
	if err != nil {
		panic(err)
	}
	for i := 0; i < 10; i++ {
		msg := fmt.Sprintf("message id is %d\n", i)
		_, err := conn.WriteMessages(
			kafka.Message{
				Value: []byte(msg),
			},
		)
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("Producer: %s", msg)
		time.Sleep(time.Second)
	}
	close(done)
}

func read() {
	conn, err := kafka.DialLeader(context.Background(), protocol, addr, topic, partition)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	for {
		msg, err := conn.ReadMessage(1024)
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("Consumer: %s", string(msg.Value))
	}
}

func main() {
	done := make(chan bool)
	go write(done)
	go read()
	<-done
}
