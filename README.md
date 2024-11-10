# Image Path Retrieval Microservice
Hi! This microservice was built to support your weather app by providing image file paths based on an imageID and a folderPath. It’s super simple to use: just send a request with an ID and folder, and it’ll respond with the exact path to your image – or let you know if there’s an issue.

## How to Use
### Sending a Request
You’ll make a POST request to /get-image-path, with a JSON payload that includes:

- **imageID (integer)**: The unique ID of the image you want.
- **folderPath (string)**: The local folder path where the images are stored.
  
#### Example POST Request (Go Code)
Here’s an example of how you’d send a request from your program. It posts an imageID and folderPath in JSON format to get back the image path.
```
go
Copy code
package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
)

func main() {
    url := "http://localhost:8080/get-image-path" // Replace with your server address if different

    // Data to send
    data := map[string]interface{}{
        "imageID":    3,
        "folderPath": "/path/to/your/images", // Replace with your actual folder path
    }

    // Encode to JSON
    jsonData, err := json.Marshal(data)
    if err != nil {
        fmt.Println("Error encoding JSON:", err)
        return
    }

    // Make the POST request
    resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
    if err != nil {
        fmt.Println("Error making request:", err)
        return
    }
    defer resp.Body.Close()
    fmt.Println("Request sent! Response status:", resp.Status)
}
```

### Getting the Response
Once you make the request, here’s what you can expect:

- If successful, you’ll get back a JSON response with filePath set to the full path of your image.
- If something went wrong (like the folder doesn’t exist, or the image ID isn’t there), you’ll get a helpful error message instead.

#### Example Response Handling (Go Code)
Here’s some code to catch the response and handle both the filePath and possible error messages.

```
go
Copy code
package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "bytes"
)

func main() {
    url := "http://localhost:8080/get-image-path"
    data := map[string]interface{}{
        "imageID":    3,
        "folderPath": "/path/to/your/images", // Update this as needed
    }
    jsonData, _ := json.Marshal(data)
    
    resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData)) 
    if err != nil {
        fmt.Println("Error making request:", err)
        return
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error reading response:", err)
        return
    }

    var result map[string]string
    if err := json.Unmarshal(body, &result); err != nil {
        fmt.Println("Error parsing JSON:", err)
        return
    }

    // Handle the response data
    if filePath, exists := result["filePath"]; exists {
        fmt.Println("Success! File Path:", filePath)
    } else if errorMsg, exists := result["error"]; exists {
        fmt.Println("Oops, Error:", errorMsg)
    }
}
```

### Example JSON Responses
Here’s what the responses from the server will look like:

#### If everything works (image found):

```
json
Copy code
{
    "filePath": "/path/to/your/images/3.png"
}
```

#### If the folder path is wrong:

```
json
Copy code
{
    "error": "Folder path does not exist. Please check the path and try again."
}
```

#### If the image ID doesn’t exist in that folder:

```
json
Copy code
{
    "error": "Image with ID 3 not found in the specified folder."
}
```

## Notes
- Folder Path: Make sure the folder path exists on your machine, and it has files named with IDs matching the ones you’re requesting.
- Server Address: If the server address isn’t http://localhost:8080, adjust accordingly in the request URL.

## UML Diagram
![image](https://github.com/user-attachments/assets/378830a2-a4be-42c1-bb2f-3dc3ce51e93e)
