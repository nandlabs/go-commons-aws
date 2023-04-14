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

func GetSession(region, bucket string) (config *aws.Config, err error) {
	if defaultSessionProvider {
		defaultSession := &provider.DefaultSession{}
		config, err = defaultSession.DefaultSessionProvider()
		return
	}
	sessionProvider := sessionProviderMap[region+bucket]
	if sessionProvider != nil {
		config, err = sessionProvider.Get()
		return
	} else {
		err = errors.New("no session provider available for region and bucket")
		return
	}
}

func AddSessionProvider(region, bucket string, provider provider.SessionProvider) {
	defaultSessionProvider = false
	sessionProviderMap[region+bucket] = provider
}
