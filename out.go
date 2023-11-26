package main

import (
	"fmt"
	"encoding/json"
	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
	_ "github.com/lib/pq" // <------------ here
)

func JsonParser(jsonData []byte) []map[string]interface{} {
	var data []map[string]interface{}
	
	

	err := json.Unmarshal(jsonData, &data)
	failOnError(err, "")
	
	resultMap := make(map[string]interface{})

	for _, item := range data {
		for key, value := range item {
			resultMap[key] = value
		}
	}

	return data
}


func main(){

	// Load .env file

	err := godotenv.Load()
	failOnError(err, "load")

	// var dbname = os.Getenv("dbname")
	// var host = os.Getenv("host")
	// var port = os.Getenv("port")
	// var user = os.Getenv("user")
	// var password = os.Getenv("password")
	// var 
	var queueName = "Test"

	// Postgres info
	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// Reading from RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	failOnError(err, "Failed to connect to RabbitMQ.")

	defer conn.Close()

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

	failOnError(err, "Cant connect to PostgreSQL.")

  	// defer db.Close()

	// err = db.Ping()

	failOnError(err, "Cant find sql server.")
	// sqlStatement := `INSERT INTO emp (ename, sal, email)`
	// Writting data to Postgres
	
	go func() {
		for msg := range msgs {
			message := JsonParser(msg.Body)
			for k,v := range message {
				fmt.Println(k)
				for k1,v1 := range v {
					fmt.Println(k1,":",v1)
				}
				failOnError(err, "Error inserting into PostgreSQL.")
			}
		}
	}()
}

func failOnError(err error, msg string) {
	if err != nil {
	  	fmt.Printf("%s: %s", msg, err)
	}
}
