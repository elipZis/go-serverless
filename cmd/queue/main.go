// Copyright elipZis GmbH 2022
// All Rights Reserved
package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/elipzis/go-serverless/queue/handler"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"os"
	"os/signal"
	"time"
)

var sess *session.Session
var svc *sqs.SQS

// Run, 8, 1
// @title Go Serverless! Example App
// @version 1.0
// @description The Example App for the Conf42: Golang 2022 Presentation "Go Serverless!"
// @contact.name Savas Ziplies
// @contact.web elipZis.com
// @contact.email goserverless@elipZis.com
func main() {
	env := os.Getenv("SERVER_ENV")
	log.Printf("Starting on %s environment", env)
	if env == "AWS" {
		lambda.Start(handler.Handler)
	} else if env == "QUEUE" {
		// Standalone/Local test code
		queueUrl := os.Getenv("QUEUE_URL")
		log.Printf("Connecting to %s", queueUrl)
		conf := aws.NewConfig().
			WithEndpoint(queueUrl).
			WithRegion(os.Getenv("AWS_REGION")).
			WithCredentials(credentials.NewStaticCredentials("123", "123", ""))
		sess = session.Must(session.NewSessionWithOptions(session.Options{
			Config: *conf,
		}))
		svc = sqs.New(sess)

		// Check for new messages every few seconds
		ticker := time.NewTicker(1 * time.Second)
		quit := make(chan os.Signal)
		go func() {
			for {
				select {
				case <-ticker.C:
					msgInput, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
						AttributeNames: []*string{
							aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
						},
						MessageAttributeNames: []*string{
							aws.String(sqs.QueueAttributeNameAll),
						},
						QueueUrl:            &queueUrl,
						MaxNumberOfMessages: aws.Int64(1),
						VisibilityTimeout:   aws.Int64(5),
					})
					if err != nil {
						log.Fatal(err)
					}

					// Iterate on all available messages
					for _, message := range msgInput.Messages {
						var msgAttributes map[string]events.SQSMessageAttribute
						msgAttributes = make(map[string]events.SQSMessageAttribute, len(message.MessageAttributes))
						for k, v := range message.MessageAttributes {
							msgAttributes[k] = events.SQSMessageAttribute{
								StringValue: v.StringValue,
								BinaryValue: v.BinaryValue,
								//StringListValues: v.StringListValues,
								BinaryListValues: v.BinaryListValues,
								DataType:         *v.DataType,
							}
						}

						//_, err = DO SOMETHING(&events.SQSMessage{
						//	MessageId:              *message.MessageId,
						//	ReceiptHandle:          *message.ReceiptHandle,
						//	Body:                   *message.Body,
						//	Md5OfBody:              *message.MD5OfBody,
						//	Md5OfMessageAttributes: *message.MD5OfMessageAttributes,
						//	MessageAttributes:      msgAttributes,
						//})

						// Clear the message from the Queue
						if err == nil {
							_, err = svc.DeleteMessage(&sqs.DeleteMessageInput{
								QueueUrl:      &queueUrl,
								ReceiptHandle: message.ReceiptHandle,
							})
						}
					}

					// In case of errors, print them
					if err != nil {
						log.Println(err)
					}

				case <-quit:
					ticker.Stop()
					return
				}
			}
		}()

		// Keep it running
		signal.Notify(quit, os.Interrupt)
		<-quit
	} else {
		// Simple local testing without Queue

	}
}
