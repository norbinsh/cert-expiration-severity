package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"

	sev "github.com/Norbinsh/cert-expiration-severity"
	"github.com/aws/aws-lambda-go/lambda"
)

type Response struct {
	StatusCode int `json:"statusCode:"`
}

func HandleRequest(ctx context.Context) (Response, error) {
	client := &http.Client{}
	target := sev.Website{"https://google.com", *client}
	cert, err := target.GetCert()
	if err != nil {
		log.Fatal(err)
	}
	severity, daysLeft := sev.GetSev(cert)

	msg := fmt.Sprintf("Severity is %d, with %d days left", severity, daysLeft)
	topicArn := os.Getenv("topic_arn")

	sess, err := session.NewSession()
	svc := sns.New(sess)

	result, err := svc.Publish(&sns.PublishInput{
		Message:  &msg,
		TopicArn: &topicArn,
	})

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(*result.MessageId)

	return Response{
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
