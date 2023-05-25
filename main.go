package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

type Ratings struct {
}

func Handler(ctx context.Context, ratings Ratings) (string, error) {
	return fmt.Sprintf("Hello World!"), nil
}

func main() {
	lambda.Start(Handler)
}
