package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	http.HandleFunc("/get-image-path", getImagePathHandler)
	fmt.Println("Microservice is running on http://localhost:8080/get-image-path")
	http.ListenAndServe(":8080", nil)
}

// getImagePathHandler handles incoming requests to find the image file path
func getImagePathHandler(w http.ResponseWriter, r *http.Request) {
	// Decode the JSON request
	var requestData struct {
		ImageID    int    `json:"imageID"`
		FolderPath string `json:"folderPath"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	fmt.Printf("Received request - Folder Path: %s, Image ID: %d\n", requestData.FolderPath, requestData.ImageID)

	// Check if the specified folder exists
	if _, err := os.Stat(requestData.FolderPath); os.IsNotExist(err) {
		fmt.Println("Error: Folder path does not exist.")
		response := map[string]string{"error": "Folder path does not exist. Please check the path and try again."}
		json.NewEncoder(w).Encode(response)
		return
	}
	fmt.Println("Folder path exists.")

	// Attempt to find the file with the given image ID
	imageFileName := strconv.Itoa(requestData.ImageID) + ".jpg" // Assuming .jpg extension for example
	imageFilePath := filepath.Join(requestData.FolderPath, imageFileName)

	// Check if the image file exists in the folder
	if _, err := os.Stat(imageFilePath); os.IsNotExist(err) {
		fmt.Printf("Error: Image with ID %d not found in the specified folder.\n", requestData.ImageID)
		response := map[string]string{"error": fmt.Sprintf("Image with ID %d not found in the specified folder.", requestData.ImageID)}
		json.NewEncoder(w).Encode(response)
		return
	}
	fmt.Printf("Image found! Returning path: %s\n", imageFilePath)

	// Send back the file path as the response
	response := map[string]string{"filePath": imageFilePath}
	json.NewEncoder(w).Encode(response)
}
