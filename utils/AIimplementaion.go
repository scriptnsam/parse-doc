package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/spf13/viper"
)

var cohereAPI string
var apiKey string

func loadConfig() error {
	API_URL := viper.GetViper().GetString("api.cohereAPI")
	API_KEY := viper.GetViper().GetString("api.cohere_key")

	cohereAPI = API_URL
	apiKey = API_KEY
	return nil

}

// AIImplementation generates documentation and writes it to a Markdown file
func AIImplementation(functionName, functionCode, outputFile string) error {
	loadConfig()
	// start spinner animation in a seperate goroutine
	done := make(chan bool)
	go func() {
		frames := []string{"|", "/", "-", "\\"} // spinner frames
		i := 0
		for {
			select {
			case <-done:
				fmt.Print("\r✅ AI Documentation Generated!    \n") // clear spinner
				return
			default:
				fmt.Printf("\r⏳ Generating API Docs %s", frames[i%len(frames)])
				i++
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	// Define payload for AI request
	payload := map[string]interface{}{
		"model": "command", // Cohere's AI Model
		"prompt": fmt.Sprintf(
			"Generate detailed API documentation (Markdown documentation) for the function below:\n\nFunction Name: %s\nCode:\n%s",
			functionName, functionCode),
		"max_tokens":  300,
		"temperature": 0.7,
	}

	// Convert payload to JSON
	jsonData, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", cohereAPI, bytes.NewBuffer(jsonData))
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	// Make request to AI API
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("❌ Error:", err)
		done <- true // stop spinner
		return err
	}
	defer resp.Body.Close()

	// Parse AI response
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	done <- true // stop spinner

	var documentation string
	if generations, ok := result["generations"].([]interface{}); ok && len(generations) > 0 {
		documentation = generations[0].(map[string]interface{})["text"].(string)
	} else {
		documentation = "No response from Cohere."
	}

	// Open or create Markdown file
	mdFile, err := os.OpenFile(outputFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("❌ Error opening Markdown file: %w", err)
	}
	defer mdFile.Close()

	// Write function documentation to Markdown
	mdContent := fmt.Sprintf(
		"### 🔹 Function: `%s`\n\n```go\n%s\n```\n\n%s\n\n---\n\n",
		functionName, functionCode, documentation,
	)
	_, err = mdFile.WriteString(mdContent)
	if err != nil {
		return fmt.Errorf("❌ Error writing to Markdown file: %w", err)
	}

	fmt.Printf("✅ Documentation for `%s` written to %s\n", functionName, outputFile)
	return nil
}
