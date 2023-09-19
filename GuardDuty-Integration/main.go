package main

import (
	"AWS-Lark-Bot/alert"
	"AWS-Lark-Bot/lib"
	"AWS-Lark-Bot/resources"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Message struct {
	Message string `json:"Message"`
}

func lambdaHandler(ctx context.Context, snsEvent events.SNSEvent) error {
	// Lark bot webhook URL
	secret := os.Getenv("WEBHOOK_KEY")
	// alert level
	alertlevel := os.Getenv("ALERT_LEVEL")
	webhookURL := "https://open.feishu.cn/open-apis/bot/v2/hook/" + secret

	// get the SNS message
	message := snsEvent.Records[0].SNS.Message

	event, err := lib.ProcessJSON(message)
	// alert message
	if err != nil {
		log.Printf("Failed to load message: %v", err)
		return err
	}
	var data resources.CardMessage
	// card message struct

	serverity := alert.GetAlertServerity(event)
	fmt.Println("serverity: ", serverity)
	lib.ProcCard(event, &data, serverity)

	threshold, err := strconv.ParseFloat(alertlevel, 64)
	if err != nil {
		return err
	}
	if serverity < threshold {
		return nil
	}
	// if serverity < threshold, don't send message to Lark Bot

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

	//body, err := io.ReadAll(resp.Body)
	//fmt.Println("resp: ", string(body))
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
