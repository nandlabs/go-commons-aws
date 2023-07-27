package sns

import (
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"go.nandlabs.io/commons-aws/provider"
	"go.nandlabs.io/commons/l3"
)

var (
	logger             = l3.Get()
	sessionProviderMap = make(map[string]provider.ConfigProvider)
	// snsClient          *sns.Client
)

func init() {
	providerSns := &ProviderSNS{}
	messagingManager := messaging.Get()
	messagingManager.Register(providerSns)
}

func GetSession(region string) (config *aws.Config, err error) {
	var p provider.ConfigProvider
	var isRegistered bool
	if p, isRegistered = sessionProviderMap[region]; !isRegistered {
		p = provider.GetDefault()
	}
	if p != nil {
		config, err = p.Get()
	} else {
		err = errors.New("no session provider available for region and bucket")
	}
	return
}

func AddSessionProvider(region string, provider provider.ConfigProvider) {
	sessionProviderMap[region] = provider
}
