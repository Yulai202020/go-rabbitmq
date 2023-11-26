package main

import (
	"fmt"
	"encoding/json"
)

func JsonParser(jsonData []byte) map[string]interface{} {
	var data []map[string]interface{}


	err := json.Unmarshal(jsonData, &data)	
	failOnError(err, "")

	return data
}

func main() {
	jsonData := []byte(`{"name":"Yulai"}`)
	a := JsonParser(jsonData)
	fmt.Println(a)
}

func failOnError(err error, msg string) {
	if err != nil {
	  	fmt.Println("%s: %s", msg, err)
	}
}