package main

import (
	"Lambda-Test/lib"
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type cloudwatchLogsEvent struct {
	AWSLogs struct {
		Data string `json:"data"`
	} `json:"awslogs"`
}

func handleRequest(ctx context.Context, event cloudwatchLogsEvent) error {
	secret := os.Getenv("WEBHOOK_KEY")
	webhookURL := "https://open.feishu.cn/open-apis/bot/v2/hook/" + secret
	// Decode from base64
	data, err := base64.StdEncoding.DecodeString(event.AWSLogs.Data)
	if err != nil {
		return fmt.Errorf("decoding error: %v", err)
	}

	// Decompress the data
	buffer := bytes.NewBuffer(data)
	reader, err := gzip.NewReader(buffer)
	if err != nil {
		return fmt.Errorf("creating gzip reader error: %v", err)
	}
	defer reader.Close()

	// Read the uncompressed data
	uncompressedData, err := io.ReadAll(reader)
	if err != nil {
		return fmt.Errorf("reading error: %v", err)
	}

	// Unmarshal the uncompressed log data
	var logs lib.Logs
	err = json.Unmarshal([]byte(uncompressedData), &logs)
	if err != nil {
		return fmt.Errorf("unmarshal log error: %v", err)
	}

	var card lib.CardMessage
	// card message struct

	lib.ProcCard(logs, &card)
	payloadBytes, err := json.Marshal(card)
	if err != nil {
		log.Printf("Failed to convert message body: %v", err)
		return fmt.Errorf("unmarshal log error: %v", err)
	}

	// POST to Lark Bot
	resp, err := http.Post(webhookURL, "application/json", strings.NewReader(string(payloadBytes)))
	if err != nil {
		log.Printf("An error occurred while sending the request to the Feishu robot: %v", err)
		return fmt.Errorf("unmarshal log error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		log.Println("The message has been successfully sent to the Feishu robot")
	} else {
		log.Printf("An error occurred while sending a message to the Feishu robot, status code: %d", resp.StatusCode)
	}

	return nil
}

func main() {
	lambda.Start(handleRequest)
}
