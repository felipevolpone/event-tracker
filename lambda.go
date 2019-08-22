package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	awslambda "github.com/aws/aws-sdk-go/service/lambda"
)

// Lambda x
type Lambda struct {
	Name      string
	Arn       string
	ReadsFrom []string
}

func allLambdas() []Lambda {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := awslambda.New(sess, &aws.Config{})
	result, _ := svc.ListFunctions(nil)

	var lambdas []Lambda

	for _, r := range result.Functions {

		newLambda := Lambda{Name: *r.FunctionArn, Arn: *r.FunctionArn}
		var subscriptions []string

		output, err := svc.ListEventSourceMappings(&awslambda.ListEventSourceMappingsInput{
			FunctionName: r.FunctionName,
		})

		if err != nil {
			fmt.Println(err)
			panic("failed")
		}

		for _, source := range output.EventSourceMappings {
			subscriptions = append(subscriptions, *source.EventSourceArn)
		}

		newLambda.ReadsFrom = subscriptions
		lambdas = append(lambdas, newLambda)
	}

	return lambdas
}

// FindLambdasThatReadEventsFromSQS x
func FindLambdasThatReadEventsFromSQS(sqs SQS) []Lambda {

	var found []Lambda
	lambdas := allLambdas()

	for i, lambda := range lambdas {
		for _, source := range lambda.ReadsFrom {
			if source == sqs.Arn {
				fmt.Println("Found it")
				found = append(found, lambdas[i])
			}
		}
	}

	return lambdas
}
