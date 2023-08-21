package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
func main() {
	LoadEnvVariables()

	r := gin.Default()

	// Authentication
	r.POST("/demo", func(c *gin.Context) {
		url := os.Getenv("SPICEDB_DOMAIN") + "/v1/schema/read"
		payload := bytes.NewBuffer([]byte("{}"))

		req, err := http.NewRequest("POST", url, payload)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Error creating request:" + err.Error()})
			return
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer secrettoken")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Error sending request:" + err.Error()})
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Error reading request:" + err.Error()})
			return
		}

		fmt.Println("Response:", string(body))
	})

	err := r.Run()
	if err != nil {
		log.Fatal("Can't start server")
	}
}
