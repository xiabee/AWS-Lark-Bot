package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"main/lib"
	"os"
)

func HandleRequest(ctx context.Context) error {
	region := os.Getenv("REGION")
	secret := os.Getenv("WEBHOOK_KEY")
	
	err := lib.SendCard(region, secret)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func main() {
	lambda.Start(HandleRequest)
}
