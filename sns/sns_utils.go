package sns

import (
	"net/url"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

func GetClient(url *url.URL) (client *sns.Client, err error) {
	err = validateMessagingUrl(url.String())
	client, err = CreateSNSClient(url)
	if err != nil {
		return
	}
	return
}

func CreateSNSClient(url *url.URL) (client *sns.Client, err error) {
	var awsSession *aws.Config

	awsSession, err = GetSession(url.Host)
	if err != nil {
		return
	}
	client = sns.NewFromConfig(*awsSession)
	return
}

func validateMessagingUrl(input string) (err error) {
	// TODO :: implementation
	return
}
