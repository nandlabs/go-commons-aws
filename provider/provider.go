package provider

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

type ConfigProvider interface {
	//Get user will create the session and send it to us for the use
	Get() (*aws.Config, error)
}

var defaultConfigProvider = &defaultProvider{}

type defaultProvider struct{}

func (d *defaultProvider) Get() (sess *aws.Config, err error) {
	var getSess aws.Config

	getSess, err = config.LoadDefaultConfig(context.Background())
	sess = &getSess
	return
}

func GetDefault() ConfigProvider {
	return defaultConfigProvider
}
