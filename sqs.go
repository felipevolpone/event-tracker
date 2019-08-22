package main

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	awssqs "github.com/aws/aws-sdk-go/service/sqs"
)

// SQS representation
type SQS struct {
	Arn           string
	SNSSubscribed []SNS
}

// Name parse name from the provided Arn
func (sqs *SQS) Name() string {
	return strings.Split(sqs.Arn, ":")[5]
}

// QueueName change the Arn to the queue format
func (sqs *SQS) QueueName() string {
	region := strings.Split(sqs.Arn, ":")[3]
	accountID := strings.Split(sqs.Arn, ":")[4]
	return fmt.Sprintf("http://sqs.%s.amazonaws.com/%s/%s", region, accountID, sqs.Name())
}

func queueDetails(sqs SQS) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	service := awssqs.New(sess)

	f := &awssqs.GetQueueAttributesInput{
		AttributeNames: []*string{aws.String("All")},
		QueueUrl:       aws.String(sqs.QueueName()),
	}

	result, err := service.GetQueueAttributes(f)
	if err != nil {
		fmt.Println("Error", err)
	}

	fmt.Println(result)
}

// FindSubscribersAndSubscriptionsFor x
func FindSubscribersAndSubscriptionsFor(sqsSubscribeds []SQS) error {

	for _, sqs := range sqsSubscribeds {
		// queueDetails(sqs)
		fmt.Println(FindLambdasThatReadEventsFromSQS(sqs))
	}

	return nil
}
