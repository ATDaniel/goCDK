package main

import (
	"fmt"
	"lambda-func/app"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Username string `json:"username"` // this is a tag
}

// Take in a payload and do something with it
func HandleRequest(event MyEvent) (string, error) {
	if event.Username == "" {
		return "", fmt.Errorf("username cannot be empty")
	}

	return fmt.Sprintf("Successfully called by - %s", event.Username), nil
}

func main() {
	myApp := app.NewApp()

	lambda.Start(myApp.ApiHandler.RegisterUserHandler) // Notice that we're not calling the function, we're passing it as a param to the AWS lambda's `Start` function
}