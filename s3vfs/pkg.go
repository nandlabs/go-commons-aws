package s3vfs

import (
	"errors"
	"net/url"
	"strings"

	"github.com/aws/aws-sdk-go/aws/session"
	"go.nandlabs.io/commons-aws/provider"
	"go.nandlabs.io/commons/vfs"
)

var (
	defaultSessionProvider = true
	sessionProviderMap     = make(map[string]provider.SessionProvider)
)

func init() {
	s3Fs := &S3Fs{}
	vfs.Register(s3Fs)
}

func GetSession(region, bucket string) (*session.Session, error) {
	if defaultSessionProvider {
		defaultSession := &provider.DefaultSession{}
		return defaultSession.Get()
	}
	sessionProvider := sessionProviderMap[region+bucket]
	if sessionProvider != nil {
		return session.Must(sessionProvider.Get()), nil
	} else {
		return nil, errors.New("no session provider available for region and bucket")
	}
}

func AddSessionProvider(region, bucket string, provider provider.SessionProvider) {
	defaultSessionProvider = false
	sessionProviderMap[region+bucket] = provider
}

func validateUrl(u *url.URL) error {
	pathElements := strings.Split(u.Path, "/")
	if len(pathElements) == 1 {
		//Only Bucket provided
		return nil
	} else if len(pathElements) >= 2 {
		//Bucket and object path provided
		return nil
	} else { //path elements==0
		//return error as it's not a valid url with bucket missing
		return errors.New("invalid url with bucket missing")
	}
}
