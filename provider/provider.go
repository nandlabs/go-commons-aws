package provider

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

type SessionProvider interface {
	//Get user will create the session and send it to us for the use
	Get() (*aws.Config, error)
}

type DefaultSession struct{}

func (defaultSession *DefaultSession) DefaultSessionProvider() (*aws.Config, error) {
	sess, err := config.LoadDefaultConfig(context.TODO())
	return &sess, err
}

func (defaultSession *DefaultSession) Get() (*aws.Config, error) {
	sess, err := config.LoadDefaultConfig(context.TODO())
	return &sess, err
}
