package sns

import (
	"context"
	"net/url"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sns/types"
)

const (
	SchemeSns = "sns"
)

var snsSchemes = []string{SchemeSns}

// ProviderSNS implements Provider Interface
type ProviderSNS struct {
	mutex sync.Mutex
}

func (snsp *ProviderSNS) Schemes() (schemes []string) {
	schemes = snsSchemes
	return
}

func (snsp *ProviderSNS) Setup() {
	// TODO : implementation
}

func (snsp *ProviderSNS) NewMessage(scheme string, options ...Option) (msg Message, err error) {
	msg = NewSNSMessage()
	return
}

func (snsp *ProviderSNS) Send(url *url.URL, msg Message, options ...Option) (err error) {
	var client *sns.Client
	// TODO client should not be created everytime
	client, err = GetClient(url)
	if err != nil {
		return
	}
	_, err = client.Publish(context.Background(), &sns.PublishInput{
		TopicArn:         aws.String("arn:aws:sns:us-east-1:*****:ok"),
		Message:          aws.String(msg),
		MessageStructure: aws.String("json"),
	})
	return
}

func (snsp *ProviderSNS) SendBatch(url *url.URL, msgs []Message, options ...Option) (err error) {
	var publishBatchEntries []types.PublishBatchRequestEntry
	var output *sns.PublishBatchOutput
	var client *sns.Client

	// TODO client should not be created everytime
	client, err = GetClient(url)
	if err != nil {
		return
	}
	for _, msg := range msgs {
		input := types.PublishBatchRequestEntry{
			Id:      aws.String("TODO"),
			Message: aws.String(msg),
		}
		publishBatchEntries = append(publishBatchEntries, input)
	}
	publishBatchInput := &sns.PublishBatchInput{
		PublishBatchRequestEntries: publishBatchEntries,
		TopicArn:                   aws.String("arn:aws:sns:us-east-1:*****:ok"),
	}
	output, err = client.PublishBatch(context.Background(), publishBatchInput)
	logger.Info(output.ResultMetadata)
	return
}

func (snsp *ProviderSNS) Receive(source *url.URL, options ...Option) (msg Message, err error) {
	return
}

func (snsp *ProviderSNS) ReceiveBatch(source *url.URL, options ...Option) (msgs []Message, err error) {
	return
}

func (snsp *ProviderSNS) AddListener(source *url.URL, listener func(msg Message), options ...Option) (err error) {
	return
}
