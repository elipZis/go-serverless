package handler

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	_ "github.com/joho/godotenv"
	"log"
)

// Handler is invoked by scraper requests from the SQS
func Handler(ctx context.Context, sqsEvent events.SQSEvent) (retVal interface{}, err error) {
	// Records should be of type ScraperRequest
	for _, message := range sqsEvent.Records {
		// DO SOMETHING
		log.Println(message)
	}

	return nil, err
}
