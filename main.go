package main

import (
	"AWS-Lark-Bot/lib"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Message struct {
	Message string `json:"Message"`
}

func lambdaHandler(ctx context.Context, snsEvent events.SNSEvent) error {
	// get the SNS message
	message := snsEvent.Records[0].SNS.Message

	// Lark bot webhook URL
	secret := os.Getenv("WEBHOOK_KEY")
	webhookURL := "https://open.feishu.cn/open-apis/bot/v2/hook/" + secret

	event, err := lib.ProcessJSON(message)
	// alert message
	if err != nil {
		log.Printf("Failed to load message: %v", err)
		return err
	}
	var data lib.CardMessage
	// card message struct

	lib.ProcCard(event, &data)
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		log.Printf("Failed to convert message body: %v", err)
		return err
	}

	// POST to Lark Bot
	resp, err := http.Post(webhookURL, "application/json", strings.NewReader(string(payloadBytes)))
	if err != nil {
		log.Printf("An error occurred while sending the request to the Feishu robot: %v", err)
		return err
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
	lambda.Start(lambdaHandler)
}
