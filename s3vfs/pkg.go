package s3vfs

import (
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"go.nandlabs.io/commons-aws/provider"
	"go.nandlabs.io/commons/l3"
	"go.nandlabs.io/commons/vfs"
)

var (
	logger                 = l3.Get()
	defaultSessionProvider = true
	sessionProviderMap     = make(map[string]provider.SessionProvider)
)

func init() {
	s3Fs := &S3Fs{}
	vfs.Register(s3Fs)
}

func GetSession(region, bucket string) (*aws.Config, error) {
	if defaultSessionProvider {
		defaultSession := &provider.DefaultSession{}
		return defaultSession.DefaultSessionProvider()
	}
	sessionProvider := sessionProviderMap[region+bucket]
	if sessionProvider != nil {
		return sessionProvider.Get()
	} else {
		return nil, errors.New("no session provider available for region and bucket")
	}
}

func AddSessionProvider(region, bucket string, provider provider.SessionProvider) {
	defaultSessionProvider = false
	sessionProviderMap[region+bucket] = provider
}
