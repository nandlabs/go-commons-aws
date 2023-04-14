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

func (defaultSession *DefaultSession) DefaultSessionProvider() (sess *aws.Config, err error) {
	var defaultSess aws.Config

	defaultSess, err = config.LoadDefaultConfig(context.Background())
	sess = &defaultSess
	return
}

func (defaultSession *DefaultSession) Get() (sess *aws.Config, err error) {
	var getSess aws.Config

	getSess, err = config.LoadDefaultConfig(context.Background())
	sess = &getSess
	return
}
