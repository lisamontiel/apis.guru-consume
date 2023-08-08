package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	apiURL := "https://api.apis.guru/v2/list.json"

	resp, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	var apiData map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&apiData)
	if err != nil {
		fmt.Println("Error JSON:", err)
		return
	}

	operations := apiData["apis"].(map[string]interface{})

	for apiName, apiInfo := range operations {
		fmt.Printf("API: %s\n", apiName)
		operationsMap := apiInfo.(map[string]interface{})
		for operationName := range operationsMap {
			fmt.Printf("  Operación: %s\n", operationName)
		}
	}
}
