package sqs

import (
	"context"
	"net/url"
	"reflect"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

const (
	SchemeSqs = "sqs"
)

var sqsSchemes = []string{SchemeSqs}

type ProviderSQS struct {
	mutex sync.Mutex
}

func (sqsp *ProviderSQS) Schemes() (schemes []string) {
	schemes = sqsSchemes
	return
}

func (sqsp *ProviderSQS) Setup() {

}

func (sqsp *ProviderSQS) NewMessage(scheme string, options ...Option) (msg Message, err error) {
	msg = NewSQSMessage()
	return
}

func (sqsp *ProviderSQS) Send(url *url.URL, msg Message, options ...Option) (err error) {
	var client *sqs.Client
	client, err = GetClient(url)
	if err != nil {
		return
	}
	// the queue URL is present as input
	sqsMsgInput := &sqs.SendMessageInput{
		MessageBody: aws.String(msg),
		QueueUrl:    aws.String(url.String()),
	}
	_, err = client.SendMessage(context.Background(), sqsMsgInput)
	return
}

func (sqsp *ProviderSQS) SendBatch(url *url.URL, msgs []Message, options ...Option) (err error) {
	var client *sqs.Client
	var publishBatchEntries []types.SendMessageBatchRequestEntry

	client, err = GetClient(url)
	if err != nil {
		return
	}

	for _, msg := range msgs {
		input := types.SendMessageBatchRequestEntry{
			// TODO :: implement uuid to send the unique id with the message
			Id:          aws.String("TODO"),
			MessageBody: aws.String(msg),
		}
		publishBatchEntries = append(publishBatchEntries, input)
	}

	publishBatchInput := &sqs.SendMessageBatchInput{
		Entries:  publishBatchEntries,
		QueueUrl: aws.String(url.String()),
	}

	_, err = client.SendMessageBatch(context.Background(), publishBatchInput)
	return
}

func (sqsp *ProviderSQS) Receive(source *url.URL, options ...Option) (msg Message, err error) {
	var client *sqs.Client

	client, err = GetClient(source)
	if err != nil {
		return
	}

	gMInput := &sqs.ReceiveMessageInput{
		MessageAttributeNames: []string{
			string(types.QueueAttributeNameAll),
		},
		QueueUrl:            aws.String(source.String()),
		MaxNumberOfMessages: 1,
		// VisibilityTimeout:   int32(*timeout),
	}
	msgResult, err := client.ReceiveMessage(context.TODO(), gMInput)
	if err != nil {
		return
	}
	msg = &MessageSQS{
		&BaseMessage{
			headers:     make(map[string]interface{}),
			headerTypes: make(map[string]reflect.Kind),
			body:        msgResult.Messages[0].Body,
		},
	}
	return
}

func (sqsp *ProviderSQS) ReceiveBatch(source *url.URL, options ...Option) (msgs []Message, err error) {
	return
}

func (sqsp *ProviderSQS) AddListener(url *url.URL, listener func(msg Message), options ...Option) (err error) {
	return
}
