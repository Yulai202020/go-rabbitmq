package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"net/http"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	failOnError(err, "Fail on read .env file.")
	
	url := os.Getenv("url")

	// JSON payload (you can customize this based on your API)
	jsonData := `[{"message":"hi"}]`

	// Create a request with the POST method and set the content type
	request, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonData)))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	request.Header.Set("Content-Type", "application/json")

	// Create an HTTP client
	client := &http.Client{}

	// Send the request
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer response.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// Print the response status code and body
	fmt.Println("Status Code:", response.Status)
	fmt.Println("Response Body:", string(body))
}

func failOnError(err error, msg string) {
	if err != nil {
	  	fmt.Printf("%s: %s", msg, err)
	}
}
