package main

import (
	"errors"
	"net/http"
	"time"
	"fmt"
	"os"
	"encoding/json"
	"github.com/streadway/amqp"
	"github.com/joho/gototenv"
)


type person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
}

func GetJson(){
	res, err := http.Get(serverHost)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
	}

	jsonData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
	}

	return jsonData
}

func main() {

	// Load .env file

	err := gototenv.Load()
	finderr(err)

	const (
		serverPort = os.Getenv("serverPort")
		serverHost = os.Getenv("serverHost")
		queueName = "Test"
	)

	// Get data to upload to RabbitMQ

	jsonData := GetJson()

	// Init

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer conn.Close()

	// pushing textMessage to rabbitMQ

	ch, err := conn.Channel()
	
	finderr(err)

	defer ch.Close()

	queue, err := ch.QueueDeclare(queueName, false, false, false, false, nil)

	finderr(err)

	fmt.Println(queue)

	err = channel.Publish(
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        jsonData,
		}
	)

	finderr(err)
}

func finderr(err any) { 
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}