// TestClient.go
package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the folder path: ")
	folderPath, _ := reader.ReadString('\n')
	folderPath = strings.TrimSpace(folderPath)

	fmt.Print("Enter the image ID: ")
	var imageID int
	_, err := fmt.Scan(&imageID)
	if err != nil {
		log.Fatal("Invalid input for image ID. Please enter a number.")
	}

	// Prepare the request data
	data := map[string]interface{}{
		"imageID":    imageID,
		"folderPath": folderPath,
	}
	jsonData, _ := json.Marshal(data)

	// Send POST request to the server
	resp, err := http.Post("http://localhost:8080/get-image-path", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal("Error sending request:", err)
	}
	defer resp.Body.Close()

	// Read and display the response
	var result map[string]string
	json.NewDecoder(resp.Body).Decode(&result)

	if filePath, exists := result["filePath"]; exists {
		fmt.Println("File Path:", filePath)
	} else if errMsg, exists := result["error"]; exists {
		fmt.Println("Error:", errMsg)
	}
}
