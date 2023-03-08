package go_commons_aws

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

var (
	AwsSession *session.Session
	err        error
)

func init() {
	AwsSession, err = session.NewSession(&aws.Config{
		Region: aws.String("REGION"),
		Credentials: credentials.NewStaticCredentials(
			"ACCESS_KEY_ID",
			"SECRET_ACCESS_KEY",
			"TOKEN"),
	})
	if err != nil {
		fmt.Println("Error creating session:", err)
		return
	}
}
