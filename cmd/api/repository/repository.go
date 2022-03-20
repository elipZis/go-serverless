package repository

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/elipzis/go-serverless/util"
	"os"
)

// Repository stores some references
type Repository struct {
	Queue *sqs.SQS
}

// NewRepository creates a repository
func NewRepository() (this *Repository) {
	this = new(Repository)
	// Connect to our required queues
	this.Queue = this.CreateQueue(os.Getenv("QUEUE_URL"))
	return this
}

// CreateQueue creates an initial connection to a queue here
func (this *Repository) CreateQueue(queueURL string) *sqs.SQS {
	//If a custom url is configured (e.g. locally)
	sess := util.NewAWSSession()
	return sqs.New(sess)
}

// SendMessage to send a simple message to the queue
func (this *Repository) SendMessage(msg []byte) (*sqs.SendMessageOutput, error) {
	queueUrl := os.Getenv("QUEUE_URL")
	return this.Queue.SendMessage(&sqs.SendMessageInput{
		DelaySeconds: aws.Int64(int64(0)),
		MessageAttributes: map[string]*sqs.MessageAttributeValue{
			"type": {
				DataType:    aws.String("String"),
				StringValue: aws.String("Message"),
			},
		},
		MessageBody: aws.String(string(msg)),
		QueueUrl:    &queueUrl,
	})
}
