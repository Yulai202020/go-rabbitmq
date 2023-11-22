package main

import (
	"os"
	"fmt"
	"github.com/streadway/amqp"
	"github.com/lib/pq"	
)

const {
	dbname = ""
	host = ""
	port = ""
	user = ""
	password = ""
	queueName = "Test"
}

const db = {
	tablename = ""
	column_name = ""
}

func main(){
	// Postgres info
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// reading from RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	failOnError(err, "Failed to connect to RabbitMQ")

	defer conn.Close()

	ch, err := conn.Channel()

	failOnError(err, "Failed to open a channel")

	defer ch.Close()

	q, err := ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	failOnError(err, "Failed to register a consumer")

	db, err := sql.Open("postgres", conndb)

	failOnError(err, "Cant connect to PostgreSQL.")

  	defer db.Close()

	// Writting data to Postgres
	
	go func() {
		for msg := range msgs {
			message := string(msg.Body)

			_, err = db.Exec("INSERT INTO $1($2) VALUES($3)", db.tablename, db.column_name, message)
			if err != nil {
				log.Println("Error inserting into PostgreSQL:", err)
			}
		}
	}()
}


func failOnError(err error, msg string) {
	if err != nil {
	  log.Panicf("%s: %s", msg, err)
	}
}	