package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, requestBody map[string]interface{}) (string, error) {
	return fmt.Sprintf("%v", requestBody), nil
}

func main() {
	lambda.Start(HandleRequest)
}
