package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/urfave/cli"
)

// SNS x
type SNS struct {
	Arn          string
	SubscribesTo []SQS
}

// SNSTrack start to track your infraesturcture based on a SNS arn
func SNSTrack(c *cli.Context) error {
	snsArn := c.Args()[0]

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := sns.New(sess)
	filter := &sns.ListSubscriptionsByTopicInput{TopicArn: &snsArn}
	result, err := svc.ListSubscriptionsByTopic(filter)

	if err != nil {
		fmt.Println("Topic not found")
		return err
	}

	var sqsSubscribeds = []SQS{}
	for _, t := range result.Subscriptions {
		sqsSubscribeds = append(sqsSubscribeds, SQS{
			Arn:           *t.Endpoint,
			SNSSubscribed: []SNS{SNS{Arn: snsArn}},
		})
	}

	fmt.Println("SQSs subscribed to the SNS:", snsArn)
	for _, topic := range sqsSubscribeds {
		fmt.Println(topic.Arn)
	}

	FindSubscribersAndSubscriptionsFor(sqsSubscribeds)

	return nil
}
