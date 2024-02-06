package main

import (
	"os"
	"fmt"
	"encoding/json"
	// "database/sql"
	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
	// _ "github.com/lib/pq" // <------------ here
)

type Table struct {
	UserId int `json:"userId"`
	ID int `json:"id"`
	Title string `json:"title"`
	Body string `json:"body"`
}

func main(){
	err := godotenv.Load()
	failOnError(err, "Fail on read .env file.")

	// var dbname = os.Getenv("dbname")
	// var host = os.Getenv("host")
	// var port = os.Getenv("port")
	// var user = os.Getenv("user")
	// var password = os.Getenv("password")
	var mode = os.Getenv("mode")
	var queueName = "Test"

	// Postgres info
	// psqlInfo := fmt.Sprintf("user=%s password=%s dbname=%s", user, password, dbname)
	
	// Connect to RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ.")
	defer conn.Close()

	
	for true {
		ch, err := conn.Channel()

		failOnError(err, "Failed to open a channel.")

		defer ch.Close()

		q, err := ch.QueueDeclare(
			queueName,
			false,
			false,
			false,
			false,
			nil,
		)
		
		failOnError(err, "Failed to declare a queue.")

		msgs, err := ch.Consume(
			q.Name,
			"",
			true,
			false,
			false,
			false,
			nil,
		)

		failOnError(err, "Failed to register a consumer.")

		// db, err := sql.Open("postgres", psqlInfo)

		// failOnError(err, "Cant connect to PostgreSQL.")

		// defer db.Close()

		// err = db.Ping()

		// failOnError(err, "Cant find sql server.")
		// // sqlStatement := `INSERT INTO emp (ename, sal, email)`
		// // Writting data to Postgres

		for msg := range msgs {
			if mode == "list" {
				var data []Table

				err := json.Unmarshal(msg.Body, &data)
				failOnError(err, "")

				fmt.Println(data)
				
				for i := range data {
					_, err = db.Exec("INSERT INTO test (userid, id, title, body) VALUES ($1, $2, $3, $4)", i.UserId, i.ID, i.Title, i.Body)	
					failOnError(err, "psql connect error.")
				}
			} else {
				var data Table

				err := json.Unmarshal(msg.Body, &data)
				failOnError(err, "")

				fmt.Println(data)

				_, err = db.Exec("INSERT INTO test (userid, id, title, body) VALUES ($1, $2, $3, $4)", data.UserId, data.ID, data.Title, data.Body)
				failOnError(err, "psql connect error.")
			}
		}
	}
}

func failOnError(err error, msg string) {
	if err != nil {
	  	fmt.Printf("%s: %s", msg, err)
	}
}
