package main

import (
	"net/http"
	"fmt"
	"os"
	"io/ioutil"
	"github.com/streadway/amqp"
	"github.com/joho/godotenv"
)

func GetJson(serverHost string) []byte {

	client := &http.Client{}
    req, err := http.NewRequest("GET", serverHost, nil)	

	failOnError(err,"error making http request.")

	resp, err := client.Do(req)

	failOnError(err, "Get json: could not read response body.")

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	failOnError(err, "Fail on convert.")

	return body
}

func main() {

	// Load .env file

	err := godotenv.Load()
	failOnError(err, "Fail on read .env file.")

	serverHost := os.Getenv("serverHost")
	const queueName = "Test"


	// Get data for site

	jsonData := GetJson(serverHost)

	// Init

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer conn.Close()

	// Pushing data to rabbitMQ

	ch, err := conn.Channel()
	
	failOnError(err, "Fail on create Channel.")

	defer ch.Close()

	queue, err := ch.QueueDeclare(queueName, false, false, false, false, nil)

	failOnError(err, "Fail on create QueueDeclare.")

	fmt.Println(queue)

	err = ch.Publish(
		"",
		queueName,
		false,
		false,
		amqp.Publishing {
			ContentType: "application/json",
			Body: []byte(jsonData),
		},
	)

	failOnError(err, "Fail on Publish.")
}

func failOnError(err error, msg string) {
	if err != nil {
	  	fmt.Println("%s: %s", msg, err)
	}
}