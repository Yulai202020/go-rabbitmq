package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
)

type PostData struct {
	Message string `json:"message"`
}

func failOnError(err error, msg string) {
	if err != nil {
	  	fmt.Println("%s: %s", msg, err)
	}
}

func Push(queueName string, jsonData []byte) {
	// Put json to rabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	failOnError(err, "Fail on connect to rabbitMQ.")

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

func main() {
	router := gin.Default()
	
	router.POST("/nonlist", func(c *gin.Context) {
		// Get json
		var postData PostData
		if err := c.ShouldBindJSON(&postData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// postData to json

		jsonData, err := json.Marshal(postData)
		failOnError(err, "")
		
		queueName := "Test"

		Push(queueName, jsonData)
	})

	router.POST("/list", func(c *gin.Context) {
		// Get json
		var postData []PostData
		if err := c.ShouldBindJSON(&postData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		
		// postData to json
		
		jsonData, err := json.Marshal(postData)
		failOnError(err, "")

		// Put json to rabbitMQ
		queueName := "Test"

		Push(queueName, jsonData)
	})

	// Run the server on localhost:8080
	err := router.Run("localhost:8080")
	if err != nil {
		panic(err)
	}
}
