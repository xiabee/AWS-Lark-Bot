package lib

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func SendCard(region string, secret string) error {

	webhookURL := "https://open.feishu.cn/open-apis/bot/v2/hook/" + secret
	var card CardMessage
	// card message struct

	ProcCard(&card, region)
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
