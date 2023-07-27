package sqs

import (
	"bytes"
	"reflect"
)

type MessageSQS struct {
	*BaseMessage
	// other properties of the sqs message
}

func NewSQSMessage() *MessageSQS {
	return &MessageSQS{
		&BaseMessage{
			headers:     make(map[string]interface{}),
			headerTypes: make(map[string]reflect.Kind),
			body:        &bytes.Buffer{},
		},
	}
}

func (lm *MessageSQS) Rsvp(yes bool, options ...Option) (err error) {
	return
}
